package controllers

import (
	"gindemo/models"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct{}

type UserApi struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

// user register

func (u UserController) Register(c *gin.Context) {
	// accept username password confirm-password
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")
	confirmPassword := c.DefaultPostForm("cpassword", "")

	if username == "" || password == "" || confirmPassword == "" {
		ReturnError(c, 4001, "Incomplete information")
		return
	}

	if password != confirmPassword {
		ReturnError(c, 4001, "Password and confirmation password are not the same")
		return
	}
	user, err := models.GetUserInfoByUsername(username)

	if user.ID != 0 {
		ReturnError(c, 4001, "Username exists")
		return
	}

	_, err = models.AddUser(username, EncryMd5(password))
	if err != nil {
		ReturnError(c, 4001, "Save user failed")
		return
	}

	ReturnSuccess(c, 200, "save success", user.ID, 1)

}

// user login

func (u UserController) Login(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	password := c.DefaultPostForm("password", "")

	if username == "" || password == "" {
		ReturnError(c, 4001, "Incomplete information")
		return

	}
	user, _ := models.GetUserInfoByUsername(username)
	if user.ID == 0 {
		ReturnError(c, 4001, "Username or password error")
		return
	}

	if user.Password != EncryMd5(password) {
		ReturnError(c, 4001, "Username or password error!")
		return
	}

	session := sessions.Default(c)
	session.Set("login:"+strconv.Itoa(int(user.ID)), user.ID)
	session.Save()

	data := UserApi{Id: int(user.ID), Username: username}
	ReturnSuccess(c, 0, "login success", data, 1)

}
