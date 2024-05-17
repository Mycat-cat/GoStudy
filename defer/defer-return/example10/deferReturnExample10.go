package main

import "fmt"

func deferRun() (res int) {
	a := 1
	b := 0

	// 测试发生panic，是否依然有deferReturn
	defer func() {
		fmt.Println(a)
	}()

	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()

	// 程序尝试对一个除数为0的数做除法，程序报panic，用defer捕获
	fmt.Println("result:", a/b)

	return a
}

func main() {
	res := deferRun()
	fmt.Println(res)
}
