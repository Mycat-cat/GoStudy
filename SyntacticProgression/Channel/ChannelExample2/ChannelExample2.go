package main

import (
	"fmt"
	"time"
)

/*
在读取channel数据的时候，用ok句式做了判断，当管道内还有数据能读取的时候，ok为true，当管道关闭且数据
读取完毕，ok 为 false。如果管道不关闭，将等待数据的输入。
*/
func main() {
	ch := make(chan int, 5)
	ch <- 1
	close(ch)
	go func() {
		for i := 0; i < 5; i++ {
			v, ok := <-ch
			if ok {
				fmt.Printf("v=%d\n", v)
			} else {
				fmt.Printf("Channel数据已读完,v=%d\n", v)
			}
		}
	}()
	time.Sleep(2 * time.Second)
}
