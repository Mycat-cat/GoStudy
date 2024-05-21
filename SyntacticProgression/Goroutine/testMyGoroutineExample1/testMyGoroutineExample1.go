package main

import "fmt"

/*
main()主协程结束的太快，我的线程可能还没执行，所以打印出来的结果大多数为end!!!
*/

func MyGoroutine() {
	fmt.Println("MyGoroutine")
}

func main() {
	go MyGoroutine()
	fmt.Println("end!!!")
}
