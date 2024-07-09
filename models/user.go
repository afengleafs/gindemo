package models

import (
	"fmt"
	"gindemo/dao"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUserTest(id int) (User, error) {
	var user User
	err := dao.InitDb().Where("id=?", id).First(&user).Error
	return user, err

}

func GetUserList() ([]User, error) {
	var users []User
	err := dao.InitDb().Find(&users).Error
	return users, err
}

func AddUser(username string) (uint, error) {
	user := User{Username: username}
	err := dao.InitDb().Create(&user).Error
	if err != nil {
		panic(err.Error())
	}
	return user.Model.ID, err
}

func UpdateUser(id uint, username string) {
	fmt.Println(id, username)
	dao.InitDb().Model(&User{}).Where("id=?", id).Update("username", username)

}

func DeleteUser(id uint) error {
	err := dao.InitDb().Where("id=?", id).Delete(&User{}).Error
	return err

}
