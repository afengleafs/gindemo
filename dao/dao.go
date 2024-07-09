package dao

import (
	"gindemo/config"
	"gindemo/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitDb() *gorm.DB {
	db, err := gorm.Open(mysql.Open(config.Mysqldb), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	if db.Error != nil {
		panic(db.Error.Error())
	}

	// 初始化数据库中的数据表
	err1 := db.AutoMigrate(&models.User{})
	if err1 != nil {
		panic(err1)
	}

	return db

}
