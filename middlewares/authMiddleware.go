package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/utils"
	"net/http"
)

func Authenticate() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")

		if _, err := utils.ParseJWT(cookie); err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
			return
		}
		c.Next()
	}
}
