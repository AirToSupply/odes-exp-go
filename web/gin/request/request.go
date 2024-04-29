package main

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/query", _query)
	router.GET("/param/user/:user_id/book/:book_id", _param)
	router.POST("/form", _form)
	router.POST("/raw", _raw)
	router.Run(":8080")
}

// 查询参数
func _query(c *gin.Context) {
	// 用法一：Query方法获取URI中携带的参数（如果没有携带参数，则返回空字符串）
	// [GET] http://localhost:8080/query?q=ok
	// [ - ] [Query] ok
	fmt.Println("[Query]", c.Query("q"))

	// 用法二：GetQuery方法获取URI中携带的参数（如果没有携带参数，则b存在性判断返回false）
	// [GET] http://localhost:8080/query?q=ok
	// [ - ] [Get Query] ok true
	// [GET] http://localhost:8080/query?q=
	// [ - ] [Get Query]  true
	// [GET] http://localhost:8080/query?q1=v
	// [ - ] [Get Query]  false
	query, b := c.GetQuery("q")
	fmt.Println("[Get Query]", query, b)

	// 用法三：GetQuery方法获取URI中携带的参数（如果没有携带参数，则b存在性判断返回false）
	// [GET] http://localhost:8080/query?q=a&q=b
	// [ - ] [Query Array] [a b], []string
	values := c.QueryArray("q")
	fmt.Printf("[Query Array] %v, %T\n", values, values)
}

// 动态参数（路径参数）
func _param(c *gin.Context) {
	// [GET] http://localhost:8080/param/user/1/book/2
	// [ - ] [Param::user_id] 1
	// [ - ] [Param::book_id] 2
	fmt.Println("[Param::user_id]", c.Param("user_id"))
	fmt.Println("[Param::book_id]", c.Param("book_id"))
}

// 表单参数
// [REQUEST]
// curl --location --request POST 'http://127.0.0.1:8080/form?id=1&hobbies=swim&name=ok' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --form 'id="1"' \
// --form 'hobbies="swim"' \
// --form 'hobbies="clamb"'
// [RESPONSE]
// [PostForm] 1
// [DefaultPostForm] x-man
// [PostFormArray] [swim clamb], [Type]: []string
// [MultipartForm] &{map[hobbies:[swim clamb] id:[1]] map[]}, [Type]: *multipart.Form
func _form(c *gin.Context) {
	// 用法一：接受单个参数
	fmt.Println("[PostForm]", c.PostForm("id"))
	// 用法二：接受单个参数（默认值）
	fmt.Println("[DefaultPostForm]", c.DefaultPostForm("name", "x-man"))
	// 用法三：接受多个参数
	hobbies := c.PostFormArray("hobbies")
	fmt.Printf("[PostFormArray] %v, [Type]: %T\n", hobbies, hobbies)
	// 用法四
	parms, _ := c.MultipartForm()
	fmt.Printf("[MultipartForm] %v, [Type]: %T\n", parms, parms)
}

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// 原始参数（Body参数）
func _raw(c *gin.Context) {
	body, _ := c.GetRawData()
	data := string(body)
	fmt.Println("[GetRawData]", string(data))
	// 解析
	var user User
	if err := json.Unmarshal(body, &user); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("[Parse JSON]", user)
}
