package service

import (
	"GOOJ/define"
	"GOOJ/helper"
	"GOOJ/models"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"
	"time"
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

// Login
// @Tags 公共方法
// @Summary 用户登录
// @Param username formData string false "username"
// @Param password formData string false "password"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /login [post]
func Login(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "username or password is empty",
		})
		return
	}

	// md5
	password = helper.GetMd5(password)

	data := new(models.UserBasic)
	err := models.DB.Where("username = ? AND password = ?", username, password).Find(&data).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "username or password is error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "login error: " + err.Error(),
		})
		return
	}

	token, err := helper.GenerateToken(data.Identity, data.Username, data.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "generate token error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"token": token,
		},
	})
}

// SendCode
// @Tags 公共方法
// @Summary 发送验证码
// @Param email formData string true "email"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /send-code [post]
func SendCode(c *gin.Context) {
	email := c.PostForm("email")
	if email == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "email is empty",
		})
		return
	}
	code := helper.GetRandomCode()
	models.RDB.Set(c, email, code, 5*time.Minute)
	err := helper.SendCode(email, code)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "send mail error: " + err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"code": code,
		},
	})
}

// Register
// @Tags 公共方法
// @Summary 用户注册
// @Param email formData string true "email"
// @Param code formData string true "code"
// @Param username formData string true "username"
// @Param password formData string true "password"
// @Param phone formData string false "phone"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /register [post]
func Register(c *gin.Context) {
	email := c.PostForm("email")
	userCode := c.PostForm("code")
	username := c.PostForm("username")
	password := c.PostForm("password")
	phone := c.PostForm("phone")

	if email == "" || userCode == "" || username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "email or userCode or username or password is empty",
		})
		return
	}

	// judge email is or not exist
	var count int64
	models.DB.Model(&models.UserBasic{}).Where("email = ?", email).Count(&count)
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "email is exist",
		})
		return
	}

	// check code
	sysCode, err := models.RDB.Get(c, email).Result()
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "get code error: " + err.Error(),
		})
		return
	}
	if sysCode != userCode {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "code is error",
		})
		return
	}

	// insert data
	data := &models.UserBasic{
		Identity: helper.GetUUID(),
		Username: username,
		Password: helper.GetMd5(password),
		Email:    email,
		Phone:    phone,
	}
	err = models.DB.Create(&data).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "register error: " + err.Error(),
		})
		return
	}

	// generate token
	token, err := helper.GenerateToken(data.Identity, data.Username, data.IsAdmin)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "generate token error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"identity": data.Identity,
			"username": data.Username,
			"token":    token,
		},
	})
}

// GetRankList
// @Tags 公共方法
// @Summary 排行榜
// @Param page formData string false "page"
// @Param size formData string false "size"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /rank-list [get]
func GetRankList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetRankList Page strconv error: ", err)
		return
	}
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))
	page = (page - 1) * size

	var count int64
	list := make([]models.UserBasic, 0)
	err = models.DB.Model(&models.UserBasic{}).Count(&count).Order("finish_problem_num DESC, submit_num ASC").Offset(page).Limit(size).Find(&list).Error
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": -1,
			"msg":  "get rank list error: " + err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": map[string]interface{}{
			"list":  list,
			"count": count,
		},
	})
}
