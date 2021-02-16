package main

import (
	"amicaldo.de/amicaldo/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	controllers.InitializeDatabase()

	r := gin.Default()
	r.Use(controllers.Cors())
	r.GET("/users", controllers.GetUsers())
	r.Run()
}
