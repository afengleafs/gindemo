package main

import (
	"gindemo/dao"
	"gindemo/models"
	"gindemo/router"
)

func main() {

	// 初始化数据库中的数据表
	err1 := dao.InitDb().AutoMigrate(&models.User{})
	if err1 != nil {
		panic(err1)
	}
	r := router.InitRouter()

	r.Run()

}
