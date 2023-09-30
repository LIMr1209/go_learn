package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func SayHello(c *gin.Context) {

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

	s, _ := c.Get("name")

	type msg struct {
		Name string // 小写 私有 无法 json 格式化 解决方法 使用 tag
		Age  int    `json:"age"`
	}

	data := msg{
		Name: s.(string),
		Age:  1,
	}
	c.JSON(http.StatusOK, data)

}

func QueryParam(c *gin.Context) {
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

func QueryForm(c *gin.Context) {
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

func QueryPath(c *gin.Context) {
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

func ShouldBind(c *gin.Context) {
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

func UploadFile(c *gin.Context) {
	//f, err := c.FormFile("file")
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"err": err.Error(),
	//	})
	//	return
	//}
	//err = c.SaveUploadedFile(f, strconv.FormatInt(int64(time.Now().Minute()), 10)+"."+f.Filename)
	//if err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{
	//		"err": err.Error(),
	//	})
	//	return
	//}
	//c.JSON(http.StatusOK, gin.H{
	//	"ok": true,
	//})

	form, _ := c.MultipartForm()
	files := form.File["file"]
	for index, file := range files {
		dst := fmt.Sprintf("./%d_%s", index, file.Filename)
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			continue
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})

}

func Redirect(c *gin.Context) {
	c.Redirect(http.StatusPermanentRedirect, "http://baidu.com")
}

func GoRoutineTest(c *gin.Context) {
	go goroutineTest(c.Copy())
	c.JSON(http.StatusOK, gin.H{
		"ok": true,
	})
}

func goroutineTest(c *gin.Context) {
	for i := 0; i < 5; i++ {
		time.Sleep(5 * time.Second)
		fmt.Println("-----------")
	}

}
