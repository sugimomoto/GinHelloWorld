package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	// HTMLを出力する際のテンプレートになる場所を指定
	engine.LoadHTMLGlob("templates/*")

	// JSON をレスポンス
	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// テンプレートを使ってHTMLをレスポンス
	engine.GET("/hello", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"insertMessage": "Hello Gin World!",
		})
	})

	// バインディングポートを指定可能
	engine.Run(":3000")
}
