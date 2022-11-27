package service

import (
	"GOOJ/models"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetUserDetail
// @Tags 公共方法
// @Summary 用户详情
// @Param user_identity query string false "user_identity"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /user-detail [get]
func GetUserDetail(c *gin.Context) {
	userIdentity := c.Query("user_identity")
	if userIdentity == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "identity is empty",
		})
		return
	}
	data := new(models.UserBasic)
	err := models.GetUserDetail(userIdentity).Find(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "get user detail error: " + err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": data,
	})
}
