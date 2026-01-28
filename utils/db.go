package utils

import (
	"gin/model"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDBConnection() *gorm.DB {
	dsn := "root:root@tcp(127.0.0.1:3306)/dev?charset=utf8mb4&parseTime=True&loc=Local"
	db, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	err := db.AutoMigrate(&model.Product{}, &model.User{}, &model.PurHistory{}, &model.Review{})
	if err != nil {
		panic("マイグレーションに失敗しました")
	}

	return db
}
