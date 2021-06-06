package server

import (
	"io"
	"os"

	"order-food-app-golang/controller"
	"order-food-app-golang/middleware"

	"github.com/gin-gonic/gin"
)

// var apiVersion string = "/v1"

// Init ...
func Init() *gin.Engine {
	r := gin.Default()
	r.RedirectTrailingSlash = true
	r.RedirectFixedPath = true
	r.RemoveExtraSlash = true

	// add loging file
	f, _ := os.Create("log/gin.log")
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	// r.Use(middleware.CORSMiddleware())

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "OK"})
	})

	users := r.Group("/user")
	userController := controller.User{}
	// users.Use(middleware.Auth())
	{
		users.GET("/", middleware.Auth(), userController.ListUser)
		users.GET("/id/:id", middleware.Auth(), userController.FindByID)
		users.POST("/login", userController.Login)
		users.POST("/register", userController.Register)
	}

	items := r.Group("/item")
	itemController := controller.Item{}
	{
		items.GET("/", middleware.Auth(), itemController.ListItem)
		items.GET("/id/:id", middleware.Auth(), itemController.FindByID)
		items.POST("/", middleware.Auth(), itemController.Create)
		items.PUT("/id/:id", middleware.Auth(), itemController.Update)
		items.PUT("/status/:id", middleware.Auth(), itemController.UpdateStatus)
	}

	orders := r.Group("/order")
	orderController := controller.Order{}
	{
		orders.GET("/", middleware.Auth(), orderController.ListOrder)
		orders.GET("/id/:id", middleware.Auth(), orderController.FindByID)
		orders.POST("/", middleware.Auth(), orderController.Create)
		orders.PUT("/id/:id", middleware.Auth(), orderController.Update)
		orders.PUT("/status/:id", middleware.Auth(), orderController.UpdateStatus)
	}

	return r
}
