package main

import (
	"amicaldo.de/amicaldo/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	controllers.InitializeDatabase()

	router := gin.Default()
	router.Use(controllers.Cors())
	users := router.Group("/users")
	{
		users.GET("/", controllers.GetUsers())
	}
	crews := router.Group("/crews")
	{
		crews.GET("/", controllers.GetCrews())
	}
	router.Run()
}
