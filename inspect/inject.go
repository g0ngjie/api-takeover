package inspect

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"takeover/file"
)

const SCRIPT_ROOT = "/inspect/scripts"

// bytesCombine 多个[]byte数组合并
func bytesCombine(pBytes ...[]byte) []byte {
	len := len(pBytes)
	s := make([][]byte, len)
	for index := 0; index < len; index++ {
		s[index] = pBytes[index]
	}
	sep := []byte("")
	return bytes.Join(s, sep)
}

// 注入调试工具
func InjectConsole(buf *bytes.Buffer, target string) {

	t := strings.ToUpper(target)
	var injectFor string

	switch t {
	case "VCONSOLE":
		injectFor = "v.js"
	case "ERUDA":
		injectFor = "ev.js"
	case "MDEBUG":
		injectFor = "mv.js"
	}

	pwd, _ := os.Getwd()
	var filePath = filepath.ToSlash(pwd + SCRIPT_ROOT + "/" + injectFor)
	b := file.ReadFile(filePath)
	const DOCTYPE = "<!DOCTYPE html>\r\n"
	top := []byte(DOCTYPE + "<script>")
	end := []byte("</script>")
	appendJs := bytesCombine(top, b, end)
	buf.Write(appendJs)
}
