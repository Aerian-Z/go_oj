package service

import (
	"go_oj/define"
	"go_oj/helper"
	"go_oj/models"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetSubmitList
// @Tags 公共方法
// @Summary 提交列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param problem_identity query string false "problem_identity"
// @Param user_identity query string false "user_identity"
// @Param status query int false "status"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /submit-list [get]
func GetSubmitList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetSubmitList Page strconv error: ", err)
		return
	}
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))

	page = (page - 1) * size
	var count int64

	problemIdentity := c.Query("problem_identity")
	userIdentity := c.Query("user_identity")
	status, _ := strconv.Atoi(c.Query("status"))

	data := make([]*models.SubmitBasic, 0)
	tx := models.GetSubmitList(problemIdentity, userIdentity, status)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("Get Submit List Error: ", err)
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

// Submit
// @Tags 用户私有方法
// @Summary 代码提交
// @Param authorization header string true "authorization"
// @Param problem_identity query string true "problem_identity"
// @Param code body string true "code"
// @Success 200 {string} json "{"code":"200","data":""}"
// @Router /user/submit [post]
func Submit(c *gin.Context) {
	problemIdentity := c.Query("problem_identity")
	code, err := io.ReadAll(c.Request.Body)
	if err != nil {
		log.Println("Submit ReadAll Error: ", err)
		return
	}

	path, err := helper.CodeSave(code)
	if err != nil {
		log.Println("Submit CodeSave Error: ", err)
		return
	}

	userClaims, _ := c.Get("user")
	data := &models.SubmitBasic{
		Identity:        helper.GetUUID(),
		ProblemIdentity: problemIdentity,
		UserIdentity:    userClaims.(*helper.UserClaims).Identity,
		Path:            path,
	}

	// judge code
	pb := new(models.ProblemBasic)
	err = models.DB.Where("identity = ?", problemIdentity).Preload("TestCases").First(pb).Error
	if err != nil {
		log.Println("Submit Get ProblemBasic Error: ", err)
		return
	}
	status := helper.JudgeCode(pb, path)
	data.Status = status

	err = models.DB.Transaction(func(tx *gorm.DB) error {
		err = tx.Create(data).Error
		if err != nil {
			log.Println("Submit Create Error: ", err)
			return err
		}
		err = tx.Model(&models.ProblemBasic{}).Where("identity = ?", problemIdentity).Update("submit_num", gorm.Expr("submit_num + ?", 1)).Error
		if err != nil {
			log.Println("Submit Update Error: ", err)
			return err
		}
		err = tx.Model(&models.UserBasic{}).Where("identity = ?", userClaims.(*helper.UserClaims).Identity).Update("submit_num", gorm.Expr("submit_num + ?", 1)).Error
		if err != nil {
			log.Println("Submit Update Error: ", err)
			return err
		}
		if status == 1 {
			err = tx.Model(&models.ProblemBasic{}).Where("identity = ?", problemIdentity).Update("pass_num", gorm.Expr("pass_num + ?", 1)).Error
			if err != nil {
				log.Println("Submit Update Error: ", err)
				return err
			}
			err = tx.Model(&models.UserBasic{}).Where("identity = ?", userClaims.(*helper.UserClaims).Identity).Update("pass_num", gorm.Expr("pass_num + ?", 1)).Error
			if err != nil {
				log.Println("Submit Update Error: ", err)
				return err
			}
		}
		return nil
	})

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"status": data.Status,
			"msg":    define.StatusMsg[data.Status],
		},
	})
}
