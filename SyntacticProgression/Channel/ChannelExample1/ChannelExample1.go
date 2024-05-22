package main

import (
	"fmt"
	"time"
)

/*
创建一个缓存为5的int类型的管道，向管道里写入一个1之后，将管道关闭，然后开启一个goroutine
从管道读取数据，读取5次，它仍然可以读取数据，在读完数据之后，将一直读取零值。
但是，上述读取方式还有一个问题?比如我们创建一个int类型的channel，我们需要往里面写入零值
用另一个goroutine读取，此时我们就无法区分读取到的是正确的零值还是数据已经读取完了而读取到
的零值。
*/
func main() {
	ch := make(chan int, 5)
	ch <- 1
	close(ch)
	go func() {
		for i := 0; i < 5; i++ {
			v := <-ch // 闭包，可以看到外面的Goroutine---ch(被捕获)
			fmt.Printf("v=%d\n", v)
		}
	}()
	time.Sleep(2 * time.Second)
}
