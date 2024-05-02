package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
	Sex  string `json:"sex"`
}

type Query struct {
	Name string `json:"name" form:"name"`
	Age  int    `json:"age" form:"age"`
	Sex  string `json:"sex" form:"sex"`
}

type Variable struct {
	UserId int `json:"uid" uri:"uid"`
	BookId int `json:"bid" uri:"bid"`
}

func main() {
	router := gin.Default()
	router.POST("/body", _body)
	router.POST("/param", _param)
	router.POST("/path/:uid/:bid", _path)
	router.POST("/form", _form)
	router.Run(":8080")
}

// ShouldBindJSON - 绑定JSON Body
// 注意：实体需要标记json tag
// <<Right>>
// [Request]
// curl --location --request POST 'http://127.0.0.1:8080/body' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw '{
// "name": "tony",
// "age": 17,
// "sex": "M"
// }'
// [Response]
//
//	   {
//		    "message": "ok",
//			"user": {
//				"name": "tony",
//				"age": 17,
//				"sex": "M"
//			}
//		}
//
// <<Wrong>>
// [Request]
// curl --location --request POST 'http://127.0.0.1:8080/body' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw '{
// "name": "tony",
// "age": "17",
// "sex": "M"
// }'
// [Response]
//
//	{
//		"error": "json: cannot unmarshal string into Go struct field User.age of type int"
//	}
func _body(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "user": user})
}

// ShouldBindQuery - 绑定查询参数
// 注意：实体需要标记form tag
// <<Right>>
// [Request]
// curl --location --request POST 'http://127.0.0.1:8080/param?name=Hobby&age=7&sex=M' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw ”
// [Response]
//
//	{
//		"message": "ok",
//		"query": {
//			"name": "Hobby",
//			"age": 7,
//			"sex": "M"
//		}
//	}
//
// <<Wrong>>
// [Request]
// curl --location --request POST 'http://127.0.0.1:8080/param?name=Hobby&age=s&sex=M' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw ”
// [Response]
//
//	{
//		"error": "strconv.ParseInt: parsing \"s\": invalid syntax"
//	}
func _param(c *gin.Context) {
	var query Query
	if err := c.ShouldBindQuery(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "query": query})
}

// ShouldBindUri - 绑定动态参数（路径参数）
// [Request]
// curl --location --request POST 'http://127.0.0.1:8080/path/1/2' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw ”
// [Response]
//
//	{
//		"message": "ok",
//		"value": {
//			"uid": 1,
//			"bid": 2
//		}
//	}
func _path(c *gin.Context) {
	var value Variable
	if err := c.ShouldBindUri(&value); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "value": value})
}

// ShouldBind - 绑定表单参数，U会根据请求头中的content-type去自动绑定，基于form-data类型的参数tag可以使用form，默认情况下也是form
// [Request]
// curl --location --request POST 'http://127.0.0.1:8080/form' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --form 'name="tom"' \
// --form 'age="19"' \
// --form 'sex="F"'
// [Response]
//
//	{
//		"message": "ok",
//		"query": {
//			"name": "tom",
//			"age": 19,
//			"sex": "F"
//		}
//	}
func _form(c *gin.Context) {
	var query Query
	if err := c.ShouldBind(&query); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "ok", "query": query})
}
