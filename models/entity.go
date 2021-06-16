package models

import "gorm.io/gorm"

// Entity ページネーションの共通化用モデル
type Entity interface {
	Count(db *gorm.DB) int64

	Take(db *gorm.DB, limit int, offset int) interface{}
}
