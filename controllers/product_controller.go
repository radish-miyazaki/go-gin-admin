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

func AllProducts(c *gin.Context) {
	// ページネーション用変数
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))

	// INFO: entityにはモデルのインスタンスを代入
	data := models.Paginate(db.DB, &models.Product{}, page)

	c.JSON(http.StatusOK, data)
}

func CreateProduct(c *gin.Context) {
	var p models.Product

	// リクエストのBodyをuに紐付ける
	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	db.DB.Create(&p)
	c.JSON(http.StatusCreated, p)
}

func GetProduct(c *gin.Context) {
	var p models.Product

	// クエリパラメータ取得
	id := c.Param("id")

	// ユーザが存在しない場合の処理
	err := db.DB.Where("id = ?", id).First(&p).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func UpdateProduct(c *gin.Context) {
	var p models.Product
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&p).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	if err := c.Bind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	db.DB.Model(&p).Updates(p)

	c.JSON(http.StatusOK, p)
}

func DeleteProduct(c *gin.Context) {
	var p models.Product
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&p).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	db.DB.Delete(&p)

	c.JSON(http.StatusNoContent, nil)
}
