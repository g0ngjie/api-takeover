package file

import (
	"os"
	"path/filepath"
	"takeover/cert"
	"takeover/util"
)

const CA_NAME = "install_ca.crt"

// 默认初始化给到用户
// 用于安装系统根证书
func init() {
	pwd, _ := os.Getwd()
	filePath := filepath.ToSlash(pwd + "/" + CA_NAME)
	if isExist, _ := PathExists(filePath); !isExist {
		f, err := os.Create(filePath)
		util.Stderr(err)
		defer f.Close()
		_, err = f.Write(cert.CA_CERT)
		util.Stderr(err)
	}
}
