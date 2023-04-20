package boot

import (
	"takeover/cert"
	"takeover/file"
)

// 顺序加载
func Initiate() {
	file.LoadYaml()
	file.LoadCfg()
	cert.SetCA()
}
