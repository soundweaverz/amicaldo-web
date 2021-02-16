package controllers

import (
	"net/http"

	. "amicaldo.de/amicaldo/models"

	"github.com/gin-gonic/gin"
)

func getCrewsFromDB() []Crew {
	rows, err := db.Query(`
SELECT name,
       image,
       group_id,
       group_id_semi_finals,
       group_id_finals,
	   playoffs,
	   disqualified,
	   customtext,
	   fpp,
	   tpp,
	   active
  FROM pubgm_crews
`)
	if err != nil {
		panic(err.Error())
	}

	crews := []Crew{}

	for rows.Next() {
		var crew Crew
		active := 0

		err = rows.Scan(&crew.Name,
			&crew.Image,
			&crew.GroupID,
			&crew.GroupIDSemiFinals,
			&crew.GroupIDFinals,
			&crew.Playoffs,
			&crew.Disqualified,
			&crew.CustomText,
			&crew.FPP,
			&crew.TPP,
			&active)
		if err != nil {
			panic(err.Error())
		}

		crew.Active = active == 1
		crews = append(crews, crew)
	}

	return crews
}

func GetCrews() func(c *gin.Context) {
	return func(c *gin.Context) {
		crews := getCrewsFromDB()
		c.JSON(http.StatusOK, crews)
	}
}
