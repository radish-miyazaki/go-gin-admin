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

func AllRoles(c *gin.Context) {
	var rs []models.Role
	db.DB.Preload("Permissions").Find(&rs)

	c.JSON(http.StatusOK, rs)
}

func CreateRole(c *gin.Context) {
	// MEMO: フロントから値を受け取る為の変数。DTO = Data Transfer Object。
	var rDTO map[string]interface{}

	if err := c.Bind(&rDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	list := rDTO["permissions"].([]interface{})

	// 受け取ったpermissionのidをPermissionインスタンスに変換
	ps := make([]models.Permission, len(list))

	for i, p := range list {
		id, _ := strconv.Atoi(p.(string))
		ps[i] = models.Permission{
			ID: uint(id),
		}
	}

	// RoleインスタンスにPermissionインスタンスをリレーション
	r := models.Role{
		Name:        rDTO["name"].(string),
		Permissions: ps,
	}

	db.DB.Create(&r)
	c.JSON(http.StatusCreated, r)
}

func GetRole(c *gin.Context) {
	var r models.Role
	id := c.Param("id")

	err := db.DB.Where("id = ?", id).Preload("Permissions").First(&r).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	c.JSON(http.StatusOK, r)
}

func UpdateRole(c *gin.Context) {
	var r models.Role
	id := c.Param("id")

	// MEMO: フロントから値を受け取る為の変数。DTO = Data Transfer Object。
	var rDTO map[string]interface{}

	if err := c.Bind(&rDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err,
		})
		return
	}

	err := db.DB.Where("id = ?", id).First(&r).Error

	if errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	list := rDTO["permissions"].([]interface{})

	// 受け取ったpermissionのidをPermissionインスタンスに変換
	ps := make([]models.Permission, len(list))
	for i, p := range list {
		id, _ := strconv.Atoi(p.(string))
		ps[i] = models.Permission{
			ID: uint(id),
		}
	}

	// TODO: 古いリレーションを削除。外部キー制約を付けていないため。
	var result interface{}
	db.DB.Table("role_permissions").Where("role_id = ?", id).Delete(&result)

	// 値を更新 + 新しいリレーションを作成
	r.Name = rDTO["name"].(string)
	r.Permissions = ps
	db.DB.Model(&r).Updates(r)

	c.JSON(http.StatusOK, r)
}

func DeleteRole(c *gin.Context) {
	var r models.Role
	id := c.Param("id")

	if err := db.DB.Where("id = ?", id).First(&r).Error; errors.Is(err, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Not found",
		})
		return
	}

	// TODO: 古いリレーションを削除。外部キー制約を付けていないため。
	var result interface{}
	db.DB.Table("role_permissions").Where("role_id = ?", id).Delete(&result)

	db.DB.Delete(&r)

	c.JSON(http.StatusNoContent, nil)
}
