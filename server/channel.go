package server

import (
	"takeover/util"

	"github.com/gin-gonic/gin"
)

func getData(c *gin.Context) {
	util.Layout(c, "ok", "成功")
}
