package main

import (
	"io"
	"os"
)

func CopyFile(dstFile, srcFile string) (wr int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	// 只有左边有一个没有声明就能用:=运算符
	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	wr, err = io.Copy(dst, src)
	dst.Close()
	src.Close()
	return
}

/*
func CopyFile(dstFile, srcFile string) (wr int64, err error) {
	src, err := os.Open(srcFile)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstFile)
	if err != nil {
		return
	}
	defer dst.Close()

	wr, err = io.Copy(dst, src)
	return wr, err
}
*/

func main() {

}
