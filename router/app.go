package router

import (
	_ "GOOJ/docs"
	"GOOJ/middlewares"
	"GOOJ/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// routers

	// public method
	// problem
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// user
	r.GET("/user-detail", service.GetUserDetail)
	r.POST("/login", service.Login)
	r.POST("/send-code", service.SendCode)
	r.POST("/register", service.Register)
	r.GET("/rank-list", service.GetRankList)
	r.GET("/category-list", service.GetCategoryList)

	// submit
	r.GET("/submit-list", service.GetSubmitList)

	// private method of administrator
	r.POST("/admin/problem-create", middlewares.AuthAdminCheck(), service.ProblemCreate)
	r.POST("/admin/category-create", middlewares.AuthAdminCheck(), service.CategoryCreate)
	r.DELETE("/admin/category-delete", middlewares.AuthAdminCheck(), service.CategoryDelete)
	r.PUT("/admin/category-modify", middlewares.AuthAdminCheck(), service.CategoryModify)

	return r
}
