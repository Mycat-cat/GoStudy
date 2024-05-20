package main

import "fmt"

func testPanic1() {
	defer func() {
		recover()
	}()
	fmt.Println("testPanic1上半部分")
	testPanic2()
	fmt.Println("testPanic1下半部分")
}

func testPanic2() {
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
