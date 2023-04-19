package util

import "fmt"

// 统一输出异常
func Stderr(err error) {
	if err != nil {
		fmt.Println("[debug]err:", err)
		panic(1)
	}
}
