package util

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ResponseLayout struct {
	Code        uint8       `json:"code"`        // 状态码
	Data        interface{} `json:"data"`        // 返回数据
	Description string      `json:"description"` // 描述
}

func getResponseBody(code uint8, data interface{}, description string) ResponseLayout {
	return ResponseLayout{
		Code:        code,
		Data:        data,
		Description: description,
	}
}

func Layout(c *gin.Context, data interface{}, description string) {
	var res = getResponseBody(100, data, description)
	c.JSON(http.StatusOK, res)
}

func ErrLayout(c *gin.Context, data interface{}, description string) {
	var res = getResponseBody(101, data, description)
	c.JSON(http.StatusOK, res)
}
