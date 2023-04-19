package file

import (
	"fmt"
	"io"
	"os"
)

// 读取到file中，再利用io将file直接读取到[]byte中
func ReadFile(path string) []byte {
	f, err := os.Open(path)
	if err != nil {
		fmt.Println("read file fail", err)
		return nil
	}
	defer f.Close()

	fd, err := io.ReadAll(f)
	if err != nil {
		fmt.Println("read to fd fail", err)
		return nil
	}

	return fd
}
