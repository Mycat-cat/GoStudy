package main

import "fmt"

/*
延迟函数	defer fmt.Println("num is :", num)的参数num在defer语句出现的时候就已经确定
num = 1,所以后面不管怎么修改a的值,最终调用defer函数传递给defer函数的参数已经固定是1了,
不会再变化
*/

func deferRun() {
	var num = 1
	//a := &num
	//defer fmt.Println("num is :", *a)
	defer fmt.Println("num is :", num)

	//*a = 4
	//a := &num
	//*a = 2
	num = 2
	return
}

func main() {
	deferRun()
}
