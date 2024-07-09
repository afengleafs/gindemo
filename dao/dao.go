package dao

import (
	"gindemo/config"
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
	return db

}
