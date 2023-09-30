package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go_leran/src/controller"
	"go_leran/src/middleware"
	"net/http"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.ForwardedByClientIP = true
	err := r.SetTrustedProxies([]string{"127.0.0.1"})
	if err != nil {
		return nil
	}
	// 告诉gin框架模板文件引用的静态文件去哪里找
	r.Static("static", "src/static")
	// 告诉gin框架去哪里找模板文件
	r.LoadHTMLGlob("src/templates/*")
	r.Use(middleware.RecordTime())
	r.GET("hello", controller.SayHello)
	r.GET("query_param", controller.QueryParam)
	r.POST("query_form", controller.QueryForm)
	r.GET("query_path/:name/:age", controller.QueryPath)

	r.POST("should_bind", controller.ShouldBind)
	r.POST("upload", controller.UploadFile)

	r.GET("redirect", controller.Redirect)
	r.GET("redirect1", func(context *gin.Context) {
		context.Request.URL.Path = "/hello"
		r.HandleContext(context)
	})

	// 匹配所有请求方法
	r.Any("/test", func(context *gin.Context) {
		fmt.Println(context.Request.URL.Path)
		switch context.Request.Method {
		case http.MethodGet:
			context.JSON(http.StatusOK, gin.H{
				"method": http.MethodGet,
			})
		case http.MethodPost:
			context.JSON(http.StatusOK, gin.H{
				"method": http.MethodGet,
			})
		}

	})

	r.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{
			"err": "Not Found",
		})
	})

	// 路由组 蓝图

	videoGroup := r.Group("video", middleware.RecordTime())
	//videoGroup.Use(middleware.RecordTime()) // 重复注册 重复执行中间件
	{
		videoGroup.GET("/list", func(context *gin.Context) {
			context.JSON(http.StatusOK, gin.H{
				"video": [3]string{"1", "2", "3"},
			})
		})
	}
	r.GET("goroutine_test", controller.GoRoutineTest)

	return r
}
