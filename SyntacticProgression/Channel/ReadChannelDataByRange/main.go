package main

import (
	"fmt"
	"time"
)

/*
在Example2中，我们明确了读取的次数是5次，但是我们往往在更多的时候，是不明确读取次数的，
只是在Channel的一端读取数据，有数据我们就读，直到另一端关闭了这个channel，
这样就可以用for range这种方式来读取channel中的数据了
*/
func main() {
	ch := make(chan int, 5)
	ch <- 1
	ch <- 2
	close(ch)
	go func() {
		for v := range ch {
			fmt.Printf("v=%d\n", v)
		}
	}()
	time.Sleep(2 * time.Second)
}

/*
这里在主goroutine关闭了channel之后，子goroutine里的for range循环才会结束。
*/
