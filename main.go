package main

import (
	"log"
	"net/http"
	"takeover/boot"
	myProxy "takeover/proxy"

	"github.com/elazarl/goproxy"
)

func init() {
	boot.Initiate()
}
func main() {

	proxy := goproxy.NewProxyHttpServer()
	proxy.Verbose = false

	println(":: 开始代理 ::")
	myProxy.CreateProxy(proxy)

	log.Fatal(http.ListenAndServe(":1234", proxy))
}
