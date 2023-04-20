package inspect

import (
	"bytes"
	"strings"
)

const SCRIPT_ROOT = "/scripts"

// bytesCombine 多个[]byte数组合并
func BytesCombine(pBytes ...[]byte) []byte {
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
	var injectFor []byte

	switch t {
	case "VCONSOLE":
		injectFor = VCONSOLE
	case "ERUDA":
		injectFor = ERUDA
	case "MDEBUG":
		injectFor = MDEBUG
	}

	buf.Write(injectFor)
}
