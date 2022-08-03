package main

import (
	mid "api-takeover/middleware"
	"flag"

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
	r.Use(mid.Collector())
	// r.StaticFS("/static", gin.Dir("tmp", true))

	r.Run(":" + port)
}
