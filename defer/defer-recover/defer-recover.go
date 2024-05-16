package main

import "fmt"

func main() {
	// 匿名函数的用法
	defer func() {
		if r := recover(); r != nil {
			fmt.Println(r)
		}
	}()
	a := 1
	b := 0
	// 程序不输出result,尝试对一个除数为0的数做除法，程序报panic，用defer捕获
	fmt.Println("result:", a/b)
}
