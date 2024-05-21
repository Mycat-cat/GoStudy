package main // 定义一个main包，它是程序的入口点。

import ( // 导入需要的包。
	"fmt"  // 用于格式化输出。
	"sync" // 提供同步原语，如WaitGroup。
	"time" // 提供时间相关的函数和类型。
)

func MyGoroutine(name string, wg *sync.WaitGroup) { // 定义一个函数MyGoroutine，接受一个字符串和一个sync.WaitGroup指针。
	defer wg.Done() // 使用defer关键字在函数返回前调用wg.Done()，表示这个goroutine已经完成。

	for i := 0; i < 5; i++ { // 一个从0到4的循环。
		fmt.Printf("MyGoroutine %s\n", name) // 打印goroutine的名称。
		time.Sleep(10 * time.Millisecond)    // 休眠10毫秒，模拟耗时操作。
	}
}

func main() { // main函数是程序执行的入口点。
	var wg sync.WaitGroup // 声明一个sync.WaitGroup变量，用于等待goroutine完成。
	wg.Add(2)             // 调用wg.Add(2)，表示需要等待两个goroutine。

	go MyGoroutine("goroutine1", &wg) // 启动第一个goroutine，并传递名称"goroutine1"和wg的地址。
	go MyGoroutine("goroutine2", &wg) // 启动第二个goroutine，并传递名称"goroutine2"和wg的地址。

	wg.Wait() // 调用wg.Wait()，阻塞主goroutine直到所有的goroutine都完成。
}
