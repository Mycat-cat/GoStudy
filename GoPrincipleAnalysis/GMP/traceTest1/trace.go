package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

func main() {
	//1. 创建一个trace文件
	f, err := os.Create("traceTest1.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	//2.启动trace
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	// 正常要调试的业务
	fmt.Println("Hello GMP")

	//3. 停止trace
	trace.Stop()
}
