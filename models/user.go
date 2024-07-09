package models

import (
	"gindemo/dao"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

// username exists？

func GetUserInfoByUsername(username string) (User, error) {
	var user User
	err := dao.InitDb().Where("username = ?", username).First(&user).Error
	return user, err
}

func AddUser(username string, password string) (int, error) {
	user := User{Username: username, Password: password}
	err := dao.InitDb().Create(&user).Error
	return int(user.ID), err
}
