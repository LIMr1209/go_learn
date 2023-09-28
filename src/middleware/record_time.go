package middleware

import "C"
import (
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func RecordTime() gin.HandlerFunc {
	return func(context *gin.Context) {
		start := time.Now()
		context.Set("name", "你好") // 请求上下文设置值
		context.Next()            // 调用剩余处理程序 后续中间件或者路由函数
		//go funcxx(context.Copy()) // goroutine 并发  中使用 context 用copy 只读副本
		//context.Abort() // 不调用剩余处理程序
		cost := time.Since(start)
		log.Printf("持续时间 %v", cost)

	}
}
