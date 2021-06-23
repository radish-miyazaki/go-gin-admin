package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email" gorm:"unique"`
	Password  []byte `json:"-"`
	RoleID    uint   `json:"role_id"`
	Role      Role   `json:"role" gorm:"foreignKey:RoleID"`
}

// SetPassword パスワードをハッシュ化して、インスタンスにセット
func (u *User) SetPassword(password string) {
	pw, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	u.Password = pw
}

// ComparePassword リクエスト値と実際の値の比較
func (u *User) ComparePassword(password string) error {
	return bcrypt.CompareHashAndPassword(u.Password, []byte(password))
}

// Count ページネーション用関数。データ数を返す。
func (u *User) Count(db *gorm.DB) int64 {
	var total int64
	db.Model(&User{}).Count(&total)

	return total
}

// Take ページネーション用関数。インスタンスのスライスを返す。
func (u *User) Take(db *gorm.DB, limit int, offset int) interface{} {
	var us []User
	db.Preload("Role").Offset(offset).Limit(limit).Find(&us)

	return us
}
