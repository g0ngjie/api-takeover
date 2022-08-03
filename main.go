package main

import (
	"flag"

	mid "api-takeover/middleware"

	"github.com/gin-gonic/gin"
)

var port string

func init() {
	flag.StringVar(&port, "p", "8080", "port")
	flag.Parse()
}

func main() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(mid.Cors())

	r.Run(":" + port)
}
