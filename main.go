package main

import (
	"log"
	"net/http"
	"takeover/boot"
	"takeover/proxy"

	"github.com/elazarl/goproxy"
)

func init() {
	boot.Initiate()
}
func main() {

	ps := goproxy.NewProxyHttpServer()
	ps.Verbose = false

	println(":: 开始代理 ::")
	proxy.CreateProxy(ps)

	log.Fatal(http.ListenAndServe(":1234", ps))
}
