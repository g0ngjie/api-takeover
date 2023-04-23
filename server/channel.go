package server

import (
	"container/list"
	"net/http"
	"takeover/util"

	"github.com/gin-gonic/gin"
)

type Channel struct {
	Status     int    `json:"status"`
	Method     string `json:"method"`
	Protocol   string `json:"protocol"`
	Host       string `json:"host"`
	Path       string `json:"path"`
	Type       string `json:"type"`
	OriginBody string `json:"originBody"`
	NewBody    string `json:"newBody"`
}

var Channels *KQueue

func init() {
	Channels = &KQueue{
		Queue: list.New(),
	}
}

// 入队
func SetData(r *http.Response, originBody, newBody string) {
	var data = Channel{
		Status:     r.StatusCode,
		Method:     r.Request.Method,
		Protocol:   r.Request.URL.Scheme,
		Host:       r.Request.Host,
		Path:       r.Request.URL.Path,
		Type:       r.Header.Get("Content-Type"),
		OriginBody: originBody,
		NewBody:    newBody,
	}
	Channels.Enqueue(data)
}

func getData(c *gin.Context) {
	var list = []Channel{}

	for Channels.Size() > 0 {
		frontCh, _ := Channels.Front()
		list = append(list, frontCh)
		Channels.Dequeue()
	}

	util.Layout(c, list, "")
}
