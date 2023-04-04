package main

import (
	"net/http"
	"context"
	"fmt"
	"io/ioutil"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

type Post struct {
    UserID       int    `json:"userId"`
    ID     int `json:"id"`
    Title string `json:"title"`
    Body    string `json:"body"`
}

func getBooks(c *gin.Context) {
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
 
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 
	}
	
	var parsedJSONObject []Post
	json.Unmarshal(body, &parsedJSONObject)

	if err != nil {
		fmt.Println(err)
		return 
	}
	c.IndentedJSON(http.StatusOK,parsedJSONObject);
}

var ctx = context.Background()

var redisInstance = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	router := gin.Default()

	err := redisInstance.Set(ctx, "topg", "Andrew Tate", 0).Err()
    if err != nil {
        panic(err)
    }

	val, err := redisInstance.Get(ctx, "ming").Result()
    if err != nil {
        panic(err)
    }

	fmt.Println(val)

	router.GET("/books", getBooks)

	router.Run(":8000")
}
