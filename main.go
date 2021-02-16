package main

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func initializeDatabase() *sql.DB {
	db, err := sql.Open("mysql", "demo_160221:demo_160221@tcp(db.amicaldo.net:3306)/demo_160221")

	if err != nil {
		panic(err.Error())
	}

	return db
}

func getUsers(db *sql.DB) []User {
	rows, err := db.Query("SELECT first_name, last_name FROM users")
	if err != nil {
		panic(err.Error())
	}

	users := []User{}

	for rows.Next() {
		var user User

		err = rows.Scan(&user.FirstName, &user.LastName)
		if err != nil {
			panic(err.Error())
		}

		users = append(users, user)
	}

	return users
}

func users(db *sql.DB) func(c *gin.Context) {
	return func(c *gin.Context) {
		users := getUsers(db)
		c.JSON(http.StatusOK, users)
	}
}

func main() {
	db := initializeDatabase()
	r := gin.Default()
	r.GET("/users", users(db))
	r.Run()
}

type User struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}
