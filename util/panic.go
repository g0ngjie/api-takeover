package util

import (
	"log"
)

// 统一输出异常
func Stderr(err error) {
	if err != nil {
		log.Panic(err)
	}
}
