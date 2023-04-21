package server

import (
	"net/http"
	"takeover/resources"

	"github.com/gin-gonic/gin"
)

func RegisterRouter() (r *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	resources.Init(r)

	//定义默认路由
	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"status": 404,
			"error":  "404, page not exists!",
		})
	})
	return
}
