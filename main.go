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

	r.GET("/login", func(c *gin.Context) {
		c.HTML(200, "login.html", gin.H{
			"message": "所有字段为必填项",
		})
	})

	r.POST("/login", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		if username == "" || password == "" {
			c.HTML(400, "login.html", gin.H{
				"message": "用户名和密码不能为空",
			})
			return
		}
		if username == "admin" && password == "123456" {
			c.HTML(200, "login.html", gin.H{
				"message": "模拟登录成功",
			})
		} else {
			c.HTML(401, "login.html", gin.H{
				"message": "用户名或密码错误",
			})
		}
	})

	r.Run() // 默认监听 :8080
}
