package file

import (
	"os"
	"path/filepath"
	"takeover/util"
)

const DIR_ROOT = "/config"
const EXAMPLE_DIR = "/api.google.com"

type FileRules struct {
	MatchUrl string
	JsonByte []byte
}

type FileMatchs struct {
	Domain   string
	RuleApis []FileRules
}

var exampleJSON = []byte(`{
    "code": "000000",
    "description": "SUCCESS!",
    "data": "okk",
    "traceId": "728685fdf08c4accb123471734f438e4.52886.16818940728383029"
}`)

var MatchRules []FileMatchs

// 示例文件
func initExampleFile(path string) {
	fileName := "api.exmaple.comabc.txt"
	filePath := filepath.ToSlash(path + "/" + fileName)
	f, err := os.Create(filePath)
	util.Stderr(err)
	defer f.Close()
	_, err = f.Write(exampleJSON)
	util.Stderr(err)
}

// 读取文件
func loadDomains(domainPath string) {
	// 读取域名文件夹
	domains, err := os.ReadDir(domainPath)
	util.Stderr(err)

	for _, domain := range domains {
		loadData := FileMatchs{
			Domain:   domain.Name(),
			RuleApis: []FileRules{},
		}

		fileRules := []FileRules{}

		// 获取接口文件
		apiPath := filepath.ToSlash(domainPath + "/" + domain.Name())
		apis, err := os.ReadDir(apiPath)
		util.Stderr(err)
		for _, api := range apis {
			getData := ReadFile(filepath.ToSlash(apiPath + "/" + api.Name()))
			// 读取文件
			fileRules = append(fileRules, FileRules{
				MatchUrl: api.Name(),
				JsonByte: getData,
			})
		}
		loadData.RuleApis = append(loadData.RuleApis, fileRules...)
		MatchRules = append(MatchRules, loadData)
	}
}

// 加载配置文件
func LoadCfg() error {

	pwd, _ := os.Getwd()
	domainPath := pwd + DIR_ROOT
	// 配置文件夹+域名文件夹
	os.MkdirAll(domainPath+EXAMPLE_DIR, os.ModePerm)
	// 示例文件
	initExampleFile(domainPath + EXAMPLE_DIR)

	loadDomains(domainPath)
	return nil
}
