package proxy

import (
	"bytes"
	"io"
	"log"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"takeover/file"
	"takeover/inspect"
	"takeover/util"

	"github.com/elazarl/goproxy"
)

// url 匹配
func realMatch(matchUrl string, realUrl *url.URL) bool {
	targetUrl := strings.Split(matchUrl, ".txt")[0]
	remoteUrl := realUrl.String()
	remoteUrl = strings.ReplaceAll(remoteUrl, "/", "")
	remoteUrl = strings.ReplaceAll(remoteUrl, "?", "")
	reg := regexp.MustCompile(targetUrl + "*")
	return reg.MatchString(remoteUrl)
}

// ssl response
func handleSSlResp(proxy *goproxy.ProxyHttpServer, sslDomain string, rule file.FileRules) {
	proxy.OnResponse(goproxy.ReqHostIs(sslDomain)).
		DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			modifyRespBody(r, rule)
			return r
		})
}

// response
func handleResp(proxy *goproxy.ProxyHttpServer, domain string, rule file.FileRules) {
	proxy.OnResponse(goproxy.ReqHostMatches(regexp.MustCompile(domain))).
		DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			modifyRespBody(r, rule)
			return r
		})
}

// modify body
func modifyRespBody(r *http.Response, rule file.FileRules) {
	if r != nil && r.Request.Method != "OPTIONS" {
		if realMatch(rule.MatchUrl, r.Request.URL) {
			log.Printf("[Modify][Body]: %s", r.Request.URL)
			r.Body = io.NopCloser(bytes.NewReader(rule.JsonByte))
		}
	}
}

// ssl script
func handleSSlScript(proxy *goproxy.ProxyHttpServer, sslDomain, target string) {
	proxy.OnResponse(goproxy.ReqHostIs(sslDomain)).
		DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			insertScript(r, target)
			return r
		})
}

// script
func handleScript(proxy *goproxy.ProxyHttpServer, domain, target string) {
	proxy.OnResponse(goproxy.ReqHostMatches(regexp.MustCompile(domain))).
		DoFunc(func(r *http.Response, ctx *goproxy.ProxyCtx) *http.Response {
			insertScript(r, target)
			return r
		})
}

func insertScript(r *http.Response, target string) {
	contentType := r.Header.Get("Content-Type")
	if strings.HasPrefix(contentType, "text/html") {
		log.Printf("[Inject][URL]: %s", r.Request.URL)
		log.Printf("[Inject][Script]: %s", target)
		buffer := bytes.NewBuffer(make([]byte, 4096))
		inspect.InjectConsole(buffer, target)
		_, err := io.Copy(buffer, r.Body)
		util.Stderr(err)
		// 读取出来后要重新写回去，如果没有最终处理，就原样返回给客户端
		r.Body = io.NopCloser(bytes.NewReader(bytes.Trim(buffer.Bytes(), "\x00")))
	}
}

// 注入脚本
func handleInject(proxy *goproxy.ProxyHttpServer, domain, target string) {
	sslDomain := domain + ":443"
	proxy.OnRequest(goproxy.ReqHostIs(sslDomain)).
		HandleConnect(goproxy.AlwaysMitm)

	handleScript(proxy, domain, target)
	handleSSlScript(proxy, sslDomain, target)
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
