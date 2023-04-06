package main

import (
        "context"
        "encoding/json"
        "fmt"
        "github.com/gin-gonic/gin"
        "github.com/redis/go-redis/v9"
        "io/ioutil"
        "net/http"
        "time"
)

type Post struct {
        UserID int    `json:"userId"`
        ID     int    `json:"id"`
        Title  string `json:"title"`
        Body   string `json:"body"`
}

func getPosts(c *gin.Context) {
        val, err := redisInstance.Get(ctx, "posts").Bytes()
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
                fmt.Println("Cache Miss")
                c.IndentedJSON(http.StatusOK, parsedJSONObject)
        } else {
                posts := []Post{}
                err = json.Unmarshal(val, &posts)
                if err != nil {
                        panic(err)
                }
                fmt.Println("Cache Hit")
                c.IndentedJSON(http.StatusOK, posts)
        }
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