package main

import "fmt"

/*
defer修饰的函数会在包含它的函数返回前执行。
执行顺序：入栈出栈顺序
defer3:参数传递按值传递,num和numPtr的值在被调用之前就已经被复制到了匿名函数的参数中,即使之后参数在外部作用域被修改,调用时使用的值仍是当初复制时候的值。
defer2:调用时捕获外部变量的值,可以看到之前和之后外部修改的值
defer1:参数直接确定,只可以看到之前外部修改的值

总结:直接记住defer传入的参数固定不变
*/

func deferRun() {
	var num = 1
	numPtr := &num
	defer fmt.Println("defer1:", numPtr, num, *numPtr) // 0xc00000a0c8 1 1
	defer func() {
		fmt.Println("defer2", numPtr, num, *numPtr) // 0xc00000a0c8 2 2
	}()
	defer func(num int, numPtr *int) {
		fmt.Println("defer3", numPtr, num, *numPtr) // 0xc00000a0c8 1 2
	}(num, numPtr)

	num = 2
	fmt.Println(numPtr, num, *numPtr) // 0xc00000a0c8 2 2
}

func main() {
	deferRun()
}
