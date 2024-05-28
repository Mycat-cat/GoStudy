package main

import (
	"fmt"
	"sync"
)

/*
程序先把wg的计数设置为10，每个for循环运行完毕都把计数器减1，main函数中执行到wg.Wait()会一直阻塞，
直到wg的计数器为零。最后打印了10个myGoroutine！是所有子goroutine任务结束后主goroutine才退出。

注意:sync.WaitGroup对象的计数器不能为负数，否则会panic，在使用的过程中，我们需要保证
add()的参数值，以及执行完Done()之后计数器大于等于零。
*/

var wg sync.WaitGroup

func myGoroutine() {
	defer wg.Done()
	fmt.Println("MyGoroutine!")
}

func main() {
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go myGoroutine()
	}
	wg.Wait()
	fmt.Println("end")
}
