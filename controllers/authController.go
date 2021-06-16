package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/models"
	"github.com/radish-miyazaki/go-admin/utils"
	"net/http"
	"strconv"
)

func Register(c *gin.Context) {
	// INFO: password_confirmというフィールドがあるので、userControllerのようにUserモデルに直接紐付けられない。
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if data["password"] != data["password_confirm"] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "passwords don't match",
		})
		return
	}

	// TODO: 登録時には確認メールを送信するユースケース追加

	u := models.User{
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
		RoleID:    1,
	}
	u.SetPassword(data["password"])
	db.DB.Create(&u)

	c.JSON(http.StatusOK, u)
}

func Login(c *gin.Context) {
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	var u models.User
	db.DB.Where("email = ?", data["email"]).First(&u)

	// 存在しないメールアドレスの場合、IDに0が入ったユーザが返ってくる
	if u.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not Found",
		})
		return
	}

	if err := u.ComparePassword(data["password"]); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Incorrect password",
		})
		return
	}

	token, err := utils.GenerateJWT(strconv.Itoa(int(u.ID)))

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.SetCookie(
		"jwt",
		token,
		60*60*24,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func Logout(c *gin.Context) {
	c.SetCookie(
		"jwt",
		"",
		-60*60*24,
		"/",
		"localhost",
		false,
		true,
	)

	c.JSON(http.StatusOK, gin.H{
		"message": "success",
	})
}

func User(c *gin.Context) {
	cookie, _ := c.Cookie("jwt")

	id, _ := utils.ParseJWT(cookie)

	var u models.User
	db.DB.Where("id = ?", id).First(&u)

	c.JSON(http.StatusOK, u)
}

func UpdateInfo(c *gin.Context) {
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	cookie, _ := c.Cookie("jwt")

	id, _ := utils.ParseJWT(cookie)
	userId, _ := strconv.Atoi(id)

	u := models.User{
		ID:        uint(userId),
		FirstName: data["first_name"],
		LastName:  data["last_name"],
		Email:     data["email"],
	}
	db.DB.Model(&u).Updates(u)

	c.JSON(http.StatusOK, u)
}

func UpdatePassword(c *gin.Context) {
	var data map[string]string
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	if data["password"] != data["password_confirm"] {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "passwords don't match",
		})
		return
	}

	cookie, _ := c.Cookie("jwt")

	id, _ := utils.ParseJWT(cookie)
	userId, _ := strconv.Atoi(id)

	u := models.User{
		ID: uint(userId),
	}
	u.SetPassword(data["password"])

	db.DB.Model(&u).Updates(u)

	c.JSON(http.StatusOK, u)
}

// TODO: パスワードリセットのfuncを追加
