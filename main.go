package main

import "C"
import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func sayHello(c *gin.Context) {

	/*data := map[string]any{
		"age":  11,
		"name": "张三",
	}

	c.JSON(http.StatusOK, data)*/

	/*var msg struct {
		Name string
		Age  int
	}
	msg.Age = 11
	msg.Name = "张三"
	c.JSON(http.StatusOK, msg)*/

	type msg struct {
		Name string // 小写 私有 无法 json 格式化 解决方法 使用 tag
		Age  int    `json:"age"`
	}

	data := msg{
		Name: "李思",
		Age:  1,
	}
	c.JSON(http.StatusOK, data)

}

func queryParam(c *gin.Context) {
	//name := c.Query("name")  // 获取get 参数 获取不到 为空字符串
	//name := c.DefaultQuery("name", "2") // 获取get 参数 获取不到 读取默认值

	//name, flag := c.GetQuery("name") //  获取get 参数 获取不到  为空字符串 flag 为 false
	//if !flag {
	//	fmt.Printf("%v", flag)
	//}

	names := c.QueryArray("name") // 获取get 参数 数组 获取不到位 nin

	c.JSON(http.StatusOK, gin.H{
		"name": names,
	})

}

func queryForm(c *gin.Context) {
	//name := c.PostForm("name") // 获取post 参数 获取不到 为空字符串
	//name := c.DefaultPostForm("name", "2") // 获取post 参数 获取不到 读取默认值

	//name, flag := c.GetPostForm("name") //  获取post 参数 获取不到  为空字符串 flag 为 false
	//if !flag {
	//	fmt.Printf("%v", flag)
	//}

	names := c.PostFormArray("name") // 获取post 参数 数组 获取不到位 nin

	c.JSON(http.StatusOK, gin.H{
		"name": names,
	})

}

func queryPath(c *gin.Context) {
	//name := c.Param("name") // 获取path 参数 获取不到 为空字符串

	names := c.Params // 获取所有path参数

	c.JSON(http.StatusOK, gin.H{
		"name": names,
	})

}

// binding:"required"` 必填

type Login struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func shouldBind(c *gin.Context) {
	var login Login
	if err := c.ShouldBind(&login); err == nil {
		fmt.Println("login success", login)
		c.JSON(http.StatusOK, login)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}

}

func main() {
	r := gin.Default() // 返回默认路由引擎
	r.GET("hello", sayHello)
	r.GET("query_param", queryParam)
	r.POST("query_form", queryForm)
	r.GET("query_path/:name/:age", queryPath)

	r.POST("should_bind", shouldBind)
	err := r.Run()
	if err != nil {
		fmt.Printf("start failed, err: %v", err)
	}

}
