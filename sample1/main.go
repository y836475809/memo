package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Sample struct {
	ID   int    `json: "id"`
	Name string `json: "name"`
}

func main() {
	router := gin.Default()

	// ルートエンドポイント
	router.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, World!",
		})
	})

	// サーバーを起動
	router.Run(":8080") // ポート8080でリッスン
}

func StartGet(c *gin.Context) {
	url := "http://foo.com/bar?a=1&b=2"
	res, err := http.Get(url)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	c.JSON(200, gin.H{"key": string(b)})
}

func StartPost(c *gin.Context) {
	sample := new(Sample)
	sample.ID = 0
	sample.Name = "hoge"
	sample_json, _ := json.Marshal(sample)
	url := "http://foo.com/post?a=1&b=10"
	res, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(sample_json))
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	b, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("err=%v\n", err)
	}
	c.JSON(200, gin.H{"key": string(b)})
}
