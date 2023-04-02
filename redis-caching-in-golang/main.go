package main

import (
	"net/http"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK,"Testing");
}

var ctx = context.Background()


func main() {
	//router setup
	router := gin.Default()

	redisDB := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Add a Key
	err := redisDB.Set(ctx, "topg", "Andrew Tate", 0).Err()
    if err != nil {
        panic(err)
    }

	//routes
	router.GET("/books", getBooks)

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
