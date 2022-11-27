package router

import (
	_ "GOOJ/docs"
	"GOOJ/service"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	// routers

	// problems
	r.GET("/problem-list", service.GetProblemList)
	r.GET("/problem-detail", service.GetProblemDetail)

	// users
	r.GET("/user-detail", service.GetUserDetail)

	// submits
	r.GET("/submit-list", service.GetSubmitList)

	return r
}
