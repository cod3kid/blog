package main

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, {data:"Hello world"})
}

func main() {
	//router setup
	router := gin.Default()

	//routes
	router.GET("/books", getBooks)

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
