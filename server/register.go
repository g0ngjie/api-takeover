package server

import "github.com/gin-gonic/gin"

func RegisterRouter() (r *gin.Engine) {
	gin.SetMode(gin.ReleaseMode)
	r = gin.Default()

	r.GET("/", hello)

	return
}
