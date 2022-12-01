package middlewares

import (
	"go_oj/helper"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthUserCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		//  check if the user is user
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
		if userClaims == nil || userClaims.IsAdmin != 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": -1,
				"msg":  "you are not user",
			})
			c.Abort()
			return
		}
		c.Set("user", userClaims)
		c.Next()
	}
}
