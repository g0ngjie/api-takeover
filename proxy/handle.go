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

	// 注入调试工具
	if file.YamlCfg.Inspect.Enabled && len(file.YamlCfg.Inspect.Domains) > 0 {
		for _, domain := range file.YamlCfg.Inspect.Domains {
			handleInject(proxy, domain, file.YamlCfg.Inspect.Target)
		}
	}
}
