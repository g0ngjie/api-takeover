package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func hello(c *gin.Context) {
	c.JSON(http.StatusOK, "ok")
}
