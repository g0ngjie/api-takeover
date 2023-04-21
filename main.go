package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"takeover/boot"
	"takeover/proxy"

	"github.com/elazarl/goproxy"
)

var (
	port string
	mode string
)

func init() {
	flag.StringVar(&port, "p", "1234", "port")
	flag.StringVar(&mode, "mode", "release", "runMode[debug, release]")
	boot.Initiate()
}

func main() {

	flag.Parse()

	ps := goproxy.NewProxyHttpServer()
	ps.Verbose = mode == "debug"

	proxy.CreateProxy(ps)

	log.Printf("server listening :%s", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), ps))
}
