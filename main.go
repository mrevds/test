package main

import (
	"ginjwt/controllers"
	"ginjwt/middlewares"
	"ginjwt/models"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	models.ConnectDataBase()
	public := r.Group("/api")

	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("api/admin")
	protected.Use(middlewares.JwtMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8081")
}
