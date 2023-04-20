package file

import (
	"os"
	"path/filepath"
	"takeover/util"

	"gopkg.in/yaml.v3"
)

type YamlInspect struct {
	Enabled bool     `yaml:"enabled"`
	Target  string   `yaml:"target"`
	Domains []string `yaml:"domains"`
}

type Yaml struct {
	Inspect YamlInspect `yaml:"inspect"`
}

const FILE_NAME = "config.yaml"

var exmapleYaml = []byte(`# # 示例配置
# inspect:
#   enabled: true
#   # vConsole|eruda|mdebug
#   target: "vConsole"
#   domains:
#     - www.google.com`)

var YamlCfg Yaml

// 加载配置
func LoadYaml() {
	pwd, _ := os.Getwd()
	filePath := filepath.ToSlash(pwd + "/" + FILE_NAME)
	if isExist, _ := PathExists(filePath); !isExist {
		f, err := os.Create(filePath)
		util.Stderr(err)
		defer f.Close()
		_, err = f.Write(exmapleYaml)
		util.Stderr(err)
	}
	dataBytes := ReadFile(filePath)

	config := Yaml{}
	err := yaml.Unmarshal(dataBytes, &config)
	util.Stderr(err)
	YamlCfg = config
}
