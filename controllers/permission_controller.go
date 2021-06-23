package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/radish-miyazaki/go-admin/db"
	"github.com/radish-miyazaki/go-admin/models"
	"net/http"
)

func AllPermissions(c *gin.Context) {
	var ps []models.Permission
	db.DB.Find(&ps)

	c.JSON(http.StatusOK, ps)
}
