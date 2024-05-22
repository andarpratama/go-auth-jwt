package main

import (
	"github.com/gin-gonic/gin"
	"go-auth-jwt/controllers"
	"go-auth-jwt/models"
)

func main(){
	models.ConnectDataBase()

	r := gin.Default()

	public := r.Group("/api")

	public.POST("/register", controllers.Register)

	r.Run(":8080")
}