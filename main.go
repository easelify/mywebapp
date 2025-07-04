package main

import (
	// 提示: 包名和目录名不一致时需要使用别名导入, 避免包名冲突或简化长包名时也可以用别名
	"github.com/easelify/mywebapp/configs/appconfig"
	"github.com/easelify/mywebapp/pkg/sqliteorm"

	"github.com/gin-gonic/gin"
)

// 定义登录表单结构体
type LoginForm struct {
	// 首字母大写的是导出字段，才能被外部访问，进行 JSON 序列化和反序列化
	// binding 标签用于验证字段，"required" 表示该字段为必填项
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

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

	// curl -X POST http://localhost:8080/login-shouldbind -H "Content-Type: application/json" -d '{"username": "Tom", "password": "25"}'
	// curl -X POST http://localhost:8080/login-shouldbind -H "Content-Type: application/json" -d '{"username": "Tom", "password": 25}'
	r.POST("/login-shouldbind", func(c *gin.Context) {
		var form LoginForm
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(400, gin.H{
				"message": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "使用 ShouldBindJSON 解析 JSON 数据成功",
			"form":    form,
		})
	})

	// curl -X POST http://localhost:8080/login-bind -H "Content-Type: application/json" -d '{"username": "Tom", "password": "25"}'
	// curl -X POST http://localhost:8080/login-bind -H "Content-Type: application/json" -d '{"username": "Tom", "password": 25}'
	r.POST("/login-bind", func(c *gin.Context) {
		var form LoginForm
		if err := c.BindJSON(&form); err != nil {
			// 返回错误已经由 BindJSON 内部处理（默认行为是直接返回 400 错误, 仅状态码，没有消息）
			// 也可以自行处理
			// c.JSON(400, gin.H{
			// 	"message": err.Error(),
			// })
			return
		}
		c.JSON(200, gin.H{
			"message": "使用 BindJSON 解析 JSON 数据成功",
			"form":    form,
		})
	})

	// 访问 /sqlite 路径时，调用 sqlite 包中的 CRUD 函数
	// curl -X GET http://localhost:8080/sqlite
	r.GET("/sqlite", func(c *gin.Context) {
		// 调用 sqlite 包中的 CRUD 函数
		sqliteorm.CRUD()
		c.JSON(200, gin.H{
			"message": "SQLite CRUD 操作已执行",
		})
	})

	r.GET("/config", func(c *gin.Context) {
		cfgs := appconfig.LoadConfig()
		c.JSON(200, gin.H{
			"configs": cfgs,
		})
	})

	r.Run() // 默认监听 :8080
}
