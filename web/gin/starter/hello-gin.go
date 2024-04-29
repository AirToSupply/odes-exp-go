package main

import (
	"github.com/gin-gonic/gin"
	_ "net/http"
)

// index
func Index(c *gin.Context) {
	c.String(200, "Hello Gin!")
}

func main() {
	// 步骤一：使用gin的Default方法创建一个路由Handler
	// router: *gin.Engine
	router := gin.Default()

	// 步骤二：通过Http方法绑定路由规则和路由函数
	// 不同于net/http库的路由函数
	// Gin进行了封装，把request和response都封装到了gin.Context的上下文环境中
	// 方式一：使用匿名函数
	router.GET("/hello", func(c *gin.Context) {
		c.String(200, "Hello World!")
	})
	// 方式二：使用显示函数
	router.GET("/index", Index)

	// 步骤三：启动路由的Run方法监听端口8080
	// 方式一：
	// router.Run本质就是http.ListenAndServe的进一步封装
	router.Run(":8080")
	// 方式二：采用原生HTTP服务的方式
	// http.ListenAndServe(":8080", router)
}
