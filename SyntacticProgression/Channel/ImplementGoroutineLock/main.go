package main

import (
	"fmt"
	"time"
)

/*
ch <- true和<- ch就相当于一个锁，将 *num =*num +1这个操作锁住了。因为ch管道的容量是1
在每个add函数里都会往channel放置一个true，直到执行完+1操作之后才将channel里的true取出。由
于channel的size是1，所以当一个goroutine在执行add函数的时候，其它goroutine执行add函数，执行
到ch <- true的时候就会阻塞，*num=*num +1不会成功，直到前一个+1操作完成，<-ch，读出了
管道的元素，这样就实现了并发安全
*/

func add(ch chan bool, num *int) {
	ch <- true
	*num = *num + 1
	<-ch
}

func main() {
	ch := make(chan bool, 1)

	var num int
	for i := 0; i < 100; i++ {
		go add(ch, &num)
	}
	time.Sleep(2)
	fmt.Println("num的值：", num)
}
