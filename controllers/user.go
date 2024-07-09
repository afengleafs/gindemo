package controllers

import (
	"gindemo/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserController struct{}

func (u UserController) GetUserInfo(c *gin.Context) {
	idStr := c.Param("id")
	name := c.Param("name")
	id, _ := strconv.Atoi(idStr)
	user, _ := models.GetUserTest(id)
	ReturnSuccess(c, 0, name, user, 1)

}

func (u UserController) AddUser(c *gin.Context) {
	username := c.DefaultPostForm("username", "")
	id, err := models.AddUser(username)
	if err != nil {
		panic(err.Error())
		return
	}
	ReturnSuccess(c, 0, "保存成功", id, 1)

}

func (u UserController) UpdateUser(c *gin.Context) {
	idStr := c.PostForm("id")
	username := c.DefaultPostForm("username", "")
	id, _ := strconv.Atoi(idStr)
	models.UpdateUser(uint(id), username)
	ReturnSuccess(c, 0, "更新成功", id, 1)
}

func (u UserController) DeleteUser(c *gin.Context) {
	idStr := c.PostForm("id")
	id, _ := strconv.Atoi(idStr)
	err := models.DeleteUser(uint(id))
	if err != nil {
		ReturnError(c, 4002, "删除错误")
	}

	ReturnSuccess(c, 0, "删除成功", true, 1)
}

func (u UserController) GetList(c *gin.Context) {

	ReturnError(c, 4004, "no information list")
}

func (u UserController) GetUserList(c *gin.Context) {
	users, err := models.GetUserList()
	if err != nil {
		ReturnError(c, 4004, "查询错误")
	}
	ReturnSuccess(c, 0, "查询成功", users, int64(len(users)))
}
