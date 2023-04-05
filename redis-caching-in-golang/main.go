package main

import (
	"net/http"
	"context"
	"fmt"
	"io/ioutil"
	"time"
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

func getPosts(c *gin.Context) {
	val, err := redisInstance.Get(ctx, "posts").Result()
    if err != nil {
	// Calling the API
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts")
 
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return 
	}
	
	// Parsing the JSON
	var parsedJSONObject []Post
	json.Unmarshal(body, &parsedJSONObject)

	// Setting the key
	redisErr := redisInstance.Set(ctx, "posts", body, 20*time.Second).Err()
    if redisErr != nil {
     	panic(redisErr)
    }
	c.IndentedJSON(http.StatusOK,parsedJSONObject);	
    }

	fmt.Printf(val)
	c.IndentedJSON(http.StatusOK, val);
}

var ctx = context.Background()

var redisInstance = redis.NewClient(&redis.Options{
	Addr: "localhost:6379",
})

func main() {
	router := gin.Default()

	router.GET("/posts", getPosts)

	router.Run(":8000")
}
