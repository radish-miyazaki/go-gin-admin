package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Upload(c *gin.Context) {
	form, err := c.MultipartForm()

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err,
		})
		return
	}

	fs := form.File["image"]
	filename := ""

	for _, f := range fs {
		filename = f.Filename
		if err := c.SaveUploadedFile(f, "./uploads/"+filename); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": err,
			})
			return
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"url": "http://localhost:8080/api/v1/uploads/" + filename,
	})
}
