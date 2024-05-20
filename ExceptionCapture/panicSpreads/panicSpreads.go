package main

import "fmt"

/*
调用链：
main->testPanic1->testPanic2->testPanic3
testPanic3中发现了panic,由于testPanic3中没有recover,向上传递
testPanic2中捕获panic,程序接着运行,由于testPanic3发生了panic,所以不再继续运行,函数跳出返回到testPanic2,testPanic2中捕获到了panic,
也不会再继续执行,跳出函数testPanic2,testPanic1接着执行
*/

func testPanic1() {
	fmt.Println("testPanic1上半部分")
	testPanic2()
	fmt.Println("testPanic1下半部分")
}

func testPanic2() {
	defer func() {
		recover()
	}()
	fmt.Println("testPanic2上半部分")
	testPanic3()
	fmt.Println("testPanic2下半部分")
}

func testPanic3() {
	fmt.Println("testPanic3上半部分")
	panic("在testPanic3出现了panic")
	fmt.Println("testPanic3下半部分")
}

func main() {
	fmt.Println("程序开始")
	testPanic1()
	fmt.Println("程序结束")
}
