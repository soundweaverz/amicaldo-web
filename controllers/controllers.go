package controllers

import (
	"database/sql"

	"github.com/gin-gonic/gin"
)

var db *sql.DB

func InitializeDatabase() {
	database, err := sql.Open("mysql", "demo_160221:demo_160221@tcp(db.amicaldo.net:3306)/demo_160221")
	if err != nil {
		panic(err.Error())
	}

	db = database
}

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")

		c.Next()
	}
}
