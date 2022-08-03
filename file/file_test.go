package file

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"testing"
	"time"
)

func Test_FileWrite(t *testing.T) {
	target, _ := os.Getwd()
	dstFile, err := os.OpenFile(target+"/test.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Fatalf("open file failed, err:%v", err)
	}
	bufWriter := bufio.NewWriter(dstFile)
	st := time.Now()
	defer func() {
		bufWriter.Flush()
		dstFile.Close()
		fmt.Println("文件写入耗时：", time.Now().Sub(st).Seconds(), "s")
	}()

	for i := 0; i < 100000; i++ {
		dstFile.Seek(0, os.SEEK_SET)
		dstFile.WriteString(strconv.Itoa(i) + "\n")

		//直接使用文件指针进行文件写操作，或者使用 bufio + 每次写操作之后flush 也能达到相同的目的(不过就没有bufio的效果了)
		//bufWriter.WriteString(strconv.Itoa(i) + "\n")
		//bufWriter.Flush()
	}
}
