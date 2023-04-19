package proxy

import (
	"takeover/file"

	"github.com/elazarl/goproxy"
)

// 创建代理
func CreateProxy(proxy *goproxy.ProxyHttpServer) {
	rules := file.MatchRules
	for _, rule := range rules {
		interceptionChannel(proxy, rule)
	}
}
