package models

import "gorm.io/gorm"

type Product struct {
	ID          uint    `json:"id"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	Image       string  `json:"image"`
	Price       float64 `json:"price"`
}

// Count ページネーション用関数。データ数を返す。
func (p *Product) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Product{}).Count(&total)

	return total
}

// Take ページネーション用関数。インスタンスのスライスを返す。
func (p *Product) Take(db *gorm.DB, limit int, offset int) interface{} {
	var ps []Product
	db.Offset(offset).Limit(limit).Find(&ps)

	return ps
}
