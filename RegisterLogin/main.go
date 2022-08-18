package main

import (
	"github.com/aom31/regislog/database"
	"github.com/aom31/regislog/handler"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {

	database.ConnectDb()
	
	r := gin.Default()
	r.POST("/register", handler.RegisterUser)
	r.Use(cors.Default())
	r.Run("localhost:8080")
}
