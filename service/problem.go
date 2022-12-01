package service

import (
	"encoding/json"
	"go_oj/define"
	"go_oj/helper"
	"go_oj/models"
	"log"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

// GetProblemList
// @Tags 公共方法
// @Summary 问题列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Param category_identity query string false "category_identity"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /problem-list [get]
func GetProblemList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv error: ", err)
		return
	}
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))

	page = (page - 1) * size

	keyword := c.Query("keyword")
	categoryIdentity := c.Query("category_identity")

	var count int64
	data := make([]*models.ProblemBasic, 0)
	tx := models.GetProblemList(keyword, categoryIdentity)
	err = tx.Count(&count).Omit("content").Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("Get Problem List Error: ", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  data,
			"count": count,
		},
	})
}

// GetProblemDetail
// @Tags 公共方法
// @Summary 问题详情
// @Param identity query string false "problem_identity"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /problem-detail [get]
func GetProblemDetail(c *gin.Context) {
	identity := c.Query("identity")
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "identity is empty",
		})
		return
	}
	data := new(models.ProblemBasic)
	tx := models.GetProblemDetail(identity)
	err := tx.First(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"data": "problem not found",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"data": "get problem detail error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}

// ProblemCreate
// @Tags 管理员私有方法
// @Summary 问题创建
// @Param authorization header string true "authorization"
// @Param title formData string true "title"
// @Param content formData string true "content"
// @Param max_runtime formData string false "max_runtime"
// @Param max_memory formData string false "max_memory"
// @Param category_ids formData array false "category_ids"
// @Param test_cases formData array true "test_cases"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /admin/problem-create [post]
func ProblemCreate(c *gin.Context) {
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMemory, _ := strconv.Atoi(c.PostForm("max_memory"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")

	// check params
	if title == "" || content == "" || len(testCases) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "params is empty",
		})
		return
	}

	data := &models.ProblemBasic{
		Identity:   helper.GetUUID(),
		Title:      title,
		Content:    content,
		MaxRuntime: maxRuntime,
		MaxMemory:  maxMemory,
	}

	problemCategories := make([]*models.ProblemCategory, 0)
	for _, id := range categoryIds {
		categoryId, _ := strconv.Atoi(id)
		problemCategories = append(problemCategories, &models.ProblemCategory{
			ProblemId:  data.ID,
			CategoryId: uint(categoryId),
		})
	}
	data.ProblemCategory = problemCategories

	// {"input":"1 2\n", "output":"3\n"}
	testCaseBasics := make([]*models.TestCase, 0)
	for _, testCase := range testCases {
		caseMap := make(map[string]string)
		err := json.Unmarshal([]byte(testCase), &caseMap)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "test case error: " + err.Error(),
			})
			return
		}
		if _, ok := caseMap["input"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "test case input is empty",
			})
			return
		}
		if _, ok := caseMap["output"]; !ok {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "test case output is empty",
			})
			return
		}
		testCaseBasic := &models.TestCase{
			Identity:        helper.GetUUID(),
			ProblemIdentity: data.Identity,
			Input:           caseMap["input"],
			Output:          caseMap["output"],
		}
		testCaseBasics = append(testCaseBasics, testCaseBasic)
	}
	data.TestCases = testCaseBasics

	// create problem
	err := models.DB.Create(data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "create problem error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
		},
	})
}

// ProblemModify
// @Tags 管理员私有方法
// @Summary 问题修改
// @Param authorization header string true "authorization"
// @Param identity query string true "identity"
// @Param title formData string false "title"
// @Param content formData string false "content"
// @Param max_runtime formData string false "max_runtime"
// @Param max_memory formData string false "max_memory"
// @Param category_ids formData array false "category_ids"
// @Param test_cases formData array false "test_cases"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /admin/problem-modify [put]
func ProblemModify(c *gin.Context) {
	identity := c.Query("identity")
	title := c.PostForm("title")
	content := c.PostForm("content")
	maxRuntime, _ := strconv.Atoi(c.PostForm("max_runtime"))
	maxMemory, _ := strconv.Atoi(c.PostForm("max_memory"))
	categoryIds := c.PostFormArray("category_ids")
	testCases := c.PostFormArray("test_cases")

	// check params
	if identity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "params is empty",
		})
		return
	}

	err := models.DB.Transaction(func(tx *gorm.DB) error {
		// update problem
		data := &models.ProblemBasic{
			Identity:   identity,
			Title:      title,
			Content:    content,
			MaxRuntime: maxRuntime,
			MaxMemory:  maxMemory,
		}

		// query problem detail
		err := tx.Model(&models.ProblemBasic{}).Where("identity = ?", identity).Find(data).Error
		if err != nil {
			return err
		}

		// update problem category
		if len(categoryIds) > 0 {
			// delete old problem category
			err = tx.Where("problem_id = ?", data.ID).Delete(&models.ProblemCategory{}).Error
			if err != nil {
				return err
			}
			// update problem category
			problemCategories := make([]*models.ProblemCategory, 0)
			for _, id := range categoryIds {
				categoryId, _ := strconv.Atoi(id)
				problemCategories = append(problemCategories, &models.ProblemCategory{
					ProblemId:  data.ID,
					CategoryId: uint(categoryId),
				})
			}
			// create new problem category
			err = tx.Create(&problemCategories).Error
			if err != nil {
				return err
			}
			data.ProblemCategory = problemCategories
		}

		if len(testCases) > 0 {
			// delete old test case
			err = tx.Where("problem_identity = ?", data.Identity).Delete(&models.TestCase{}).Error
			if err != nil {
				return err
			}
			// {"input":"1 2\n", "output":"3\n"}
			testCaseBasics := make([]*models.TestCase, 0)
			for _, testCase := range testCases {
				caseMap := make(map[string]string)
				err = json.Unmarshal([]byte(testCase), &caseMap)
				if err != nil {
					return err
				}
				if _, ok := caseMap["input"]; !ok {
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "test case input is empty",
					})
					return nil
				}
				if _, ok := caseMap["output"]; !ok {
					c.JSON(http.StatusOK, gin.H{
						"code": -1,
						"msg":  "test case output is empty",
					})
					return nil
				}
				testCaseBasic := &models.TestCase{
					Identity:        helper.GetUUID(),
					ProblemIdentity: data.Identity,
					Input:           caseMap["input"],
					Output:          caseMap["output"],
				}
				testCaseBasics = append(testCaseBasics, testCaseBasic)

				// create new test case
				err = tx.Create(&testCaseBasic).Error
				if err != nil {
					return err
				}
				data.TestCases = testCaseBasics
			}
		}

		// update problem
		return tx.Model(&models.ProblemBasic{}).Where("identity = ?", identity).Updates(data).Error
	})
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "modify problem error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": identity,
		},
	})
}
