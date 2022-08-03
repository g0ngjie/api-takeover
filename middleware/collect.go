package middleware

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// 接口收集
func Collector() gin.HandlerFunc {
	return func(c *gin.Context) {

		// 获取请求的url
		url := c.Request.URL.String()

		fmt.Println("[debug]url:", url)
		fmt.Println("[debug]c.Request:", c.Request)
	}
}
