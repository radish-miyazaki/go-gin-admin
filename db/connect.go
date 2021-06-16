package db

import (
	"github.com/radish-miyazaki/go-admin/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	// DBの接続チェック
	connect, err := gorm.Open(mysql.Open("root:password@/go_admin"), &gorm.Config{})
	if err != nil {
		panic("couldn't connect to the database")
	}

	// 他のモジュールで使えるように切り出し
	DB = connect

	// TODO: gormではなく、専用のmigrationライブラリを探す
	// TODO: AutoMigrateの後で、AddForeignKeyをチェーンさせることで外部キー制約を設定
	err = connect.AutoMigrate(&models.User{}, &models.Role{}, &models.Permission{},
		&models.Product{}, &models.Order{}, &models.OrderItem{})
	if err != nil {
		panic("couldn't migrate model to database")
	}
}
