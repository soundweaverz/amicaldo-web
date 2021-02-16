package controllers

import (
	"net/http"

	. "amicaldo.de/amicaldo/models"

	"github.com/gin-gonic/gin"
)

func getUsersFromDB() []User {
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

func GetUsers() func(c *gin.Context) {
	return func(c *gin.Context) {
		users := getUsersFromDB()
		c.JSON(http.StatusOK, users)
	}
}
