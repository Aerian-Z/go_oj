package service

import (
	"go_oj/define"
	"go_oj/helper"
	"go_oj/models"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetCategoryList
// @Tags 公共方法
// @Summary 分类列表
// @Param page query int false "page"
// @Param size query int false "size"
// @Param keyword query string false "keyword"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /category-list [get]
func GetCategoryList(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", define.DefaultPage))
	if err != nil {
		log.Println("GetProblemList Page strconv error: ", err)
		return
	}
	size, _ := strconv.Atoi(c.DefaultQuery("size", define.DefaultSize))

	page = (page - 1) * size

	keyword := c.Query("keyword")
	var count int64

	data := make([]*models.CategoryBasic, 0)
	tx := models.GetCategoryList(keyword)
	err = tx.Count(&count).Offset(page).Limit(size).Find(&data).Error
	if err != nil {
		log.Println("Get Category List Error: ", err)
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

// CategoryCreate
// @Tags 管理员私有方法
// @Summary 分类创建
// @Param authorization header string true "authorization"
// @Param name formData string true "name"
// @Param parent_id formData int false "parent_id"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /admin/category-create [post]
func CategoryCreate(c *gin.Context) {
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	category := &models.CategoryBasic{
		Identity: helper.GetUUID(),
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Create(&category).Error
	if err != nil {
		log.Println("Create Category Error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "create category success",
	})
}

// CategoryDelete
// @Tags 管理员私有方法
// @Summary 分类删除
// @Param authorization header string true "authorization"
// @Param identity query string true "identity"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /admin/category-delete [delete]
func CategoryDelete(c *gin.Context) {
	identity := c.Query("identity")
	var count int64
	err := models.DB.Where("category_id = (SELECT id FROM category_basic WHERE identity = ? LIMIT 1)", identity).
		Count(&count).Error
	if count > 0 {
		c.JSON(http.StatusOK, gin.H{
			"code": 200,
			"msg":  "category has problem",
		})
		return
	}

	err = models.DB.Where("identity = ?", identity).Delete(&models.CategoryBasic{}).Error
	if err != nil {
		log.Println("Delete Category Error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "delete category success",
	})
}

// CategoryModify
// @Tags 管理员私有方法
// @Summary 分类修改
// @Param authorization header string true "authorization"
// @Param identity query string true "identity"
// @Param name formData string false "name"
// @Param parent_id formData int false "parent_id"
// @Success 200 {string} json "{"code":"200","msg":""}"
// @Router /admin/category-modify [put]
func CategoryModify(c *gin.Context) {
	identity := c.Query("identity")
	name := c.PostForm("name")
	parentId, _ := strconv.Atoi(c.PostForm("parent_id"))
	category := &models.CategoryBasic{
		Identity: identity,
		Name:     name,
		ParentId: parentId,
	}
	err := models.DB.Model(new(models.CategoryBasic)).Where("identity = ?", identity).Updates(&category).Error
	if err != nil {
		log.Println("Modify Category Error: ", err)
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "modify category success",
	})
}
