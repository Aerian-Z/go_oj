package middlewares

import (
	"GOOJ/helper"
	"github.com/gin-gonic/gin"
	"net/http"
)

func AuthAdminCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  check if the user is administrator
		auth := c.GetHeader("Authorization")
		userClaims, err := helper.AnalysisToken(auth)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "token analysis error: " + err.Error(),
			})
			c.Abort()
			return
		}
		if userClaims == nil || userClaims.IsAdmin != 1 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "you are not admin",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
