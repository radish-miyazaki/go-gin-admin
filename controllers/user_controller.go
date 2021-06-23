package controllers

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/models"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

func AllUsers(c *gin.Context) {
	c.Set("page", "users")
	// ページネーション用変数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// INFO: entityにはモデルのインスタンスを代入
	data := models.Paginate(db.DB, &models.User{}, page)

	c.JSON(http.StatusOK, data)
}

func CreateUser(c *gin.Context) {
	var u models.User

	// リクエストのBodyをuに紐付ける
	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	// 管理側でのユーザ作成は固定パスワード
	u.SetPassword("password")
	db.DB.Create(&u)
	c.JSON(http.StatusCreated, u)
}

func GetUser(c *gin.Context) {
	var u models.User

	// クエリパラメータ取得
	id := c.Param("id")

	// ユーザが存在しない場合の処理
	err := db.DB.Where("id = ?", id).Preload("Role").First(&u).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, u)
}

func UpdateUser(c *gin.Context) {
	var u models.User
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&u).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	if err := c.Bind(&u); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	db.DB.Model(&u).Updates(u)

	c.JSON(http.StatusOK, u)
}

func DeleteUser(c *gin.Context) {
	var u models.User
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&u).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	db.DB.Delete(&u)

	c.JSON(http.StatusNoContent, nil)
}
