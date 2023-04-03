package main

import (
	"net/http"
	"context"
	"fmt"
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

	redisInstance := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	// Setting a Key
	err := redisInstance.Set(ctx, "topg", "Andrew Tate", 0).Err()
    if err != nil {
        panic(err)
    }
	// Getting a Key
	val, err := redisInstance.Get(ctx, "ming").Result()
    if err != nil {
        panic(err)
    }

	fmt.Println(val)
	//routes
	router.GET("/books", getBooks)

	// Listening to port..  by default it listens to 8080 router.Run()
	router.Run(":8000")
}
