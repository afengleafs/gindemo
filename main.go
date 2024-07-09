package main

import (
	"gindemo/dao"
	"gindemo/models"
	"gindemo/router"
)

func main() {

	r := router.InitRouter()
	db := dao.InitDb()
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		panic(err)
	}
	r.Run()

}
