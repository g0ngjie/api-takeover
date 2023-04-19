package proxy

import (
	"bytes"
	"io"
	"net/http"
	"regexp"
	"strings"
	"takeover/file"

	"github.com/elazarl/goproxy"
)

// 获取真实match url
func realMatch(matchUrl string) string {
	targetUrl := strings.Split(matchUrl, ".txt")[0]
	targetUrl = strings.ReplaceAll(targetUrl, "_", "/")
	return "^/" + targetUrl + "*"
}

// ssl响应
func handleSSlResp(proxy *goproxy.ProxyHttpServer, sslDomain string, rule file.FileRules) {
	proxy.OnResponse(goproxy.ReqHostIs(sslDomain)).
		DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			if r != nil {
				if r.Request.Method == "OPTIONS" {
					return r
				}

				matchPath := realMatch(rule.MatchUrl)
				reg := regexp.MustCompile(matchPath)
				if reg.MatchString(r.Request.URL.Path) {
					r.Body = io.NopCloser(bytes.NewReader(rule.JsonByte))
					return r
				}
			}
			return r
		})
}

// 响应
func handleResp(proxy *goproxy.ProxyHttpServer, domain string, rule file.FileRules) {
	proxy.OnResponse(goproxy.ReqHostMatches(regexp.MustCompile(domain))).
		DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {

			matchPath := realMatch(rule.MatchUrl)
			reg := regexp.MustCompile(matchPath)
			if reg.MatchString(r.Request.URL.Path) {
				r.Body = io.NopCloser(bytes.NewReader(rule.JsonByte))
			}
			return r
		})
}

// 拦截通道
func interceptionChannel(proxy *goproxy.ProxyHttpServer, rule file.FileMatchs) {
	sslDomain := rule.Domain + ":443"
	proxy.OnRequest(goproxy.ReqHostIs(sslDomain)).
		HandleConnect(goproxy.AlwaysMitm)

	for _, api := range rule.RuleApis {
		handleResp(proxy, rule.Domain, api)
		handleSSlResp(proxy, sslDomain, api)
	}
}
