package main

import (
	"database/sql"
	"log"
	"net/http"
	"time"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gin-gonic/gin"
)

type Attraction struct {
	Id         string ` db:"id"  json:"id"`
	Name       string `db:"name" json:"name"`
	Detail     string `db:"detail"	json:"detail"`
	Coverimage string `db:"coverimage"	json:"coverimage"`
	Latitude   string `db:"latitude"	json:"latitude"`
	Longtitude string `db:"longtitude" 	json:"longtitude"`
}

var db *sql.DB

func main() {
	var err error
	//connect db
	db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/gosqlapi")
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	db.SetConnMaxLifetime(time.Minute * 3)
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(10)

	router := gin.Default()
	router.GET("/attractions", getAttraction)

	router.Run("localhost:8080")
}

//when call api response json
func getAttraction(c *gin.Context) {
	// var (
	// 	id         int
	// 	name       string
	// 	detail     string
	// 	coverimage string
	// 	latitude   float64
	// 	longtitude float64
	// )

	//from struct
	var attractions []Attraction
	//query data
	rows, err := db.Query("select * from attractions ")
	//check
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close() //call wheen func finish
	for rows.Next() {
		var a Attraction //recive data each rows
		// err := rows.Scan(&id, &name, &detail, &coverimage, &latitude, &longtitude)
		err := rows.Scan(a.Id, a.Name, a.Detail, a.Coverimage, a.Latitude, a.Longtitude)
		if err != nil {
			log.Fatal(err)
		}
		attractions = append(attractions, a)

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	c.IndentedJSON(http.StatusOK, attractions)
}
