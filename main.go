package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("web/template/*.html")

	r.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"message": "Hello, World!",
		})
	})

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.Run() // 默认监听 :8080
}
