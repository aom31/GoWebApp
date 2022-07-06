package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Artist string `json:"artist"`
	Price  string `json:"price"`
}

//create album data array
var albums = []album{
	{ID: "1", Title: "BNK", Artist: "bnk48", Price: "2100"},
	{ID: "2", Title: "Boy", Artist: "pakorn", Price: "1200"},
	{ID: "3", Title: "bank", Artist: "cash", Price: "1400"},
	{ID: "4", Title: "singto", Artist: "numto", Price: "2500"},
}

func main() {
	fmt.Println("hello go")
	router := gin.Default()
	router.GET("/albums", getAlbum)

	router.Run("localhost:8080")
}

//when call api response json
func getAlbum(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, albums)
}
