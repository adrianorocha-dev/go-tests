package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type album struct {
	Id     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float64 `json:"price"`
}

var albums = []album{
	{Id: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
	{Id: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
	{Id: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
}

func getAlbums(ctx *gin.Context) {
	ctx.IndentedJSON(http.StatusOK, albums)
}

func postAlbum(ctx *gin.Context) {
	var newAlbum album

	if error := ctx.BindJSON(&newAlbum); error != nil {
		return
	}

	albums = append(albums, newAlbum)

	ctx.IndentedJSON(http.StatusCreated, newAlbum)
}

func getAlbumById(ctx *gin.Context) {
	id := ctx.Param("id")

	for _, a := range albums {
		if a.Id == id {
			ctx.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	ctx.IndentedJSON(http.StatusNotFound, gin.H{"message": "Album not found"})
}

func main() {
	router := gin.Default()

	router.GET("/albums", getAlbums)
	router.GET("/albums/:id", getAlbumById)
	router.POST("/albums", postAlbum)

	router.Run("localhost:8080")
}
