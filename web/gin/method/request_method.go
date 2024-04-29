package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Article struct {
	Id     int    `json:"id"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

type Response struct {
	Code    int    `json:"code"`
	Data    any    `json:"data"`
	Message string `json:"message"`
}

func main() {
	router := gin.Default()
	router.GET("/articles", GetList)
	router.GET("/articles/:id", GetDetail)
	router.POST("/articles", Create)
	router.PUT("/articles/:id", Update)
	router.DELETE("/articles/:id", Delete)
	router.Run(":8080")
}

// GetList 文章列表
// curl --location --request GET 'http://127.0.0.1:8080/articles' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive'
func GetList(c *gin.Context) {
	articles := []Article{
		Article{
			Id:     1,
			Title:  "《水浒传》",
			Author: "施耐庵",
		},
		Article{
			Id:     2,
			Title:  "《三国演义》",
			Author: "罗贯中",
		},
		Article{
			Id:     3,
			Title:  "《西游记》",
			Author: "吴承恩",
		},
		Article{
			Id:     4,
			Title:  "《红楼梦》",
			Author: "曹雪芹",
		},
	}
	c.JSON(http.StatusOK, &Response{0, articles, "success"})
}

// GetDetail 文章详情
// curl --location --request GET 'http://127.0.0.1:8080/articles/1' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw ”
func GetDetail(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "valid article id"})
		return
	}
	article := Article{
		Id:     id,
		Title:  "《水浒传》",
		Author: "施耐庵",
	}
	c.JSON(http.StatusOK, &Response{0, article, "success"})
}

// Create 添加文章
// curl --location --request POST 'http://127.0.0.1:8080/articles' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw '{"title": "《桃花源记》", "author": "陶渊明"}'
func Create(c *gin.Context) {
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "valid request value"})
		return
	}
	var article Article
	if err = json.Unmarshal(data, &article); err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "parse request failed"})
		return
	}
	c.JSON(http.StatusOK, &Response{0, article, "success"})
}

// Update 编辑文章
// curl --location --request PUT 'http://127.0.0.1:8080/articles/1' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Content-Type: application/json' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive' \
// --data-raw '{"title": "《桃花源记》", "author": "陶渊明"}'
func Update(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "valid article id"})
		return
	}
	data, err := c.GetRawData()
	if err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "valid request value"})
		return
	}
	var article Article
	if err = json.Unmarshal(data, &article); err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "parse request failed"})
		return
	}
	article.Id = id
	c.JSON(http.StatusOK, &Response{0, article, "success"})
}

// Delete 删除文章
// curl --location --request DELETE 'http://127.0.0.1:8080/articles/1' \
// --header 'User-Agent: Apifox/1.0.0 (https://apifox.com)' \
// --header 'Accept: */*' \
// --header 'Host: 127.0.0.1:8080' \
// --header 'Connection: keep-alive'
func Delete(c *gin.Context) {
	if _, err := strconv.Atoi(c.Param("id")); err != nil {
		c.JSON(http.StatusInternalServerError, &Response{500, nil, "valid article id"})
		return
	}
	c.JSON(http.StatusOK, &Response{0, nil, "success"})
}
