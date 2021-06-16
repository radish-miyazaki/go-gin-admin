package models

import "gorm.io/gorm"

type Order struct {
	ID         uint        `json:"id"`
	FirstName  string      `json:"-"`
	LastName   string      `json:"-"`
	Name       string      `json:"name" gorm:"-"`  // INFO: DBには保存されない属性
	Total      float32     `json:"total" gorm:"-"` // INFO: DBには保存されない属性
	Email      string      `json:"email"`
	UpdatedAt  string      `json:"updated_at"`
	CreatedAt  string      `json:"created_at"`
	OrderItems []OrderItem `json:"order_items" gorm:"foreignKey:OrderID"`
}

// Count ページネーション用関数。データ数を返す。
func (p *Order) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&Order{}).Count(&total)

	return total
}

// Take ページネーション用関数。インスタンスのスライスを返す。
func (p *Order) Take(db *gorm.DB, limit int, offset int) interface{} {
	var os []Order
	db.Preload("OrderItems").Offset(offset).Limit(limit).Find(&os)

	// DBには存在しない属性(Name, Total)をセット
	for i, _ := range os {
		var total float32

		for _, oi := range os[i].OrderItems {
			total += oi.Price * float32(oi.Quantity)
		}

		os[i].Name = os[i].FirstName + " " + os[i].LastName
		os[i].Total = total
	}

	return os
}
