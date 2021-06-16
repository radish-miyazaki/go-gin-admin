package models

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"math"
)

func Paginate(db *gorm.DB, entity Entity, page int) gin.H {
	// 1ページあたりのデータ数
	limit := 15
	// どのテータから取るかのオフセット
	offset := (page - 1) * limit

	// 各々のモデルで定義されているInterfaceの関数呼び出し
	total := entity.Count(db)
	data := entity.Take(db, limit, offset)

	return gin.H{
		"data": data,
		"meta": gin.H{
			"total":     total,
			"page":      page,
			"last_page": math.Ceil(float64(total) / float64(limit)), // Ceilで切り上げ
		},
	}
}
