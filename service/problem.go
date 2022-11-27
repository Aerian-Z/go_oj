package service

import (
	"GOOJ/define"
	"GOOJ/models"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

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
