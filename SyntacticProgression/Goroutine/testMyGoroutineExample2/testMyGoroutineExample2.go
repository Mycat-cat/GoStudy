package main

import (
	"fmt"
	"time"
)

/*
这次能输出了
*/

func MyGoroutine() {
	fmt.Println("MyGoroutine")
}

func main() {
	go MyGoroutine()
	fmt.Println("end!!!")
	time.Sleep(2 * time.Second)
}
