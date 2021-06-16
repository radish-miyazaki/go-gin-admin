package middlewares

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/models"
	"github.com/radish-miyazaki/go-admin/utils"
	"net/http"
	"strconv"
	"strings"
)

func Authorized() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookie, _ := c.Cookie("jwt")
		id, err := utils.ParseJWT(cookie)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "Unauthorized",
			})
			c.Abort()
		}

		userId, _ := strconv.Atoi(id)

		// ログイン中のユーザーインスタンス取得
		u := models.User{
			ID: uint(userId),
		}
		db.DB.Preload("Role").Find(&u)

		// ユーザーからロールインスタンスを取得
		r := models.Role{
			ID: u.RoleID,
		}
		db.DB.Preload("Permissions").Find(&r)

		// FIXME: ページ名をURLから取得しているので他の方法を模索。
		method := c.Request.Method
		url := c.Request.URL.Path           // ex.) /api/v1/users
		urlSlice := strings.Split(url, "/") // ex.) [ ,api ,v1 ,users]
		page := urlSlice[3]                 // ex.) users
		fmt.Println(method, url, page)

		if method == "GET" {
			for _, p := range r.Permissions {
				if p.Name == "edit_"+page || p.Name == "view_"+page {
					c.Next()
					return
				}
			}
		} else {
			for _, p := range r.Permissions {
				if p.Name == "edit_"+page {
					c.Next()
					return
				}
			}
		}

		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "No Permission",
		})
		c.Abort()
	}
}
