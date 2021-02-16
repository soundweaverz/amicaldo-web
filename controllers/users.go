package controllers

import (
	"net/http"

	. "amicaldo.de/amicaldo/models"

	"github.com/gin-gonic/gin"
)

func getUsersFromDB() []User {
	rows, err := db.Query("SELECT name, ingame_id, image, active FROM pubgm_users")
	if err != nil {
		panic(err.Error())
	}

	users := []User{}

	for rows.Next() {
		var user User
		active := 0

		err = rows.Scan(&user.Name, &user.IngameID, &user.Image, &active)
		if err != nil {
			panic(err.Error())
		}

		user.Active = active == 1
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
