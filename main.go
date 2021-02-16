package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

var db = InitializeDatabase()

func InitializeDatabase() *sql.DB {
	db, err := sql.Open("mysql", "demo_160221:demo_160221@tcp(db.amicaldo.net:3306)/demo_160221")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run()
}
