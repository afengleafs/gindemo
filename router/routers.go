package router

import (
	"gindemo/controllers"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "hello world")
	})

	user := r.Group("/user")
	{
		user.GET("/info/:id", controllers.UserController{}.GetUserInfo)

		user.POST("/add", controllers.UserController{}.AddUser)

		user.POST("/update", controllers.UserController{}.UpdateUser)

		user.GET("/list", controllers.UserController{}.GetList)

		user.GET("/all", controllers.UserController{}.GetUserList)

		user.POST("/delete", controllers.UserController{}.DeleteUser)
	}

	order := r.Group("/order")
	{
		order.POST("/list", controllers.OrderController{}.GetList)
	}

	return r

}
