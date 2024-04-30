package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// 用户实体
// json Tag用于将返回的Key指定为定义的名称，例如：例子中姓名，默认情况下不指定返回JOSN的key为Name，指定之后为uname
// json Tag标记为"-"表示不返回该字段
type User struct {
	// 姓名
	Name string `json:"uname"`
	// 年龄
	Age int `json:"age"`
	// 密码
	Password string `json:"-"`
}

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("web/gin/template/*")
	router.GET("/string", _string)
	router.GET("/json", _json)
	router.GET("/xml", _xml)
	router.GET("/yaml", _yaml)
	router.GET("/html", _html)
	router.GET("/redirect", _redirect)
	router.GET("/header", _header)
	router.Run(":8080")
}

// 响应：字符串
func _string(c *gin.Context) {
	c.String(http.StatusOK, "Hello Gin!")
}

// 响应：JSON
func _json(c *gin.Context) {
	// 用法一
	c.JSON(http.StatusOK, &User{Name: "Leo", Age: 13, Password: "123456"})
	// 用法二
	// c.JSON(http.StatusOK, map[string]interface{}{"Name": "Alex", "Age": 18})
	// 用法三
	// c.JSON(http.StatusOK, gin.H{"Name": "Sean", "Age": 19})
}

// 响应：XML
func _xml(c *gin.Context) {
	c.XML(http.StatusOK, gin.H{"Name": "Sean", "Age": 19})
}

// 响应：YAML
// 在浏览器方式下，是直接下载文件
func _yaml(c *gin.Context) {
	c.YAML(http.StatusOK, gin.H{"Name": "Sean", "Age": 19})
}

// 响应：HTML
// 需要通过LoadHTMLGlob方法加载HTML文件模版
func _html(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{"username": "Gina"})
}

// 重定向
func _redirect(c *gin.Context) {
	// 重定向域名 state code: 301 or 302
	// c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	// 重定向站内路由
	c.Redirect(http.StatusFound, "/html")
}

// 设置响应头
func _header(c *gin.Context) {
	c.Header("Token", "87d7yd74r")
}
