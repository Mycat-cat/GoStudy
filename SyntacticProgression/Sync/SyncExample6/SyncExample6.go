package main

import (
	"fmt"
	"sync"
	"time"
)

/*
简单分析:首先goroutine3开始加了读锁，开始读取，读到count的值为0，然后goroutine1尝试写
入，goroutine2尝试写入，但是都会阻塞，因为goroutine3加了读锁，不能再加写锁，在第8行
goroutine3 读取完毕之后，goroutine1争抢到了锁，加了写锁，写完释放写锁之后，goroutine1和
goroutine2同时加了读锁，读到count的值为1。可以看到读写锁是互斥的，写写锁是互斥的，读读锁
可以一起加。
*/

var cnt = 0

func main() {
	var mr = sync.RWMutex{}
	for i := 1; i <= 3; i++ {
		go write(&mr, i)
	}
	for i := 1; i <= 3; i++ {
		go read(&mr, i)
	}

	time.Sleep(time.Second)
	fmt.Println("final count:", cnt)
}

func read(mr *sync.RWMutex, i int) {
	fmt.Printf("goroutine%d reader start\n", i)
	mr.RLock()
	fmt.Printf("goroutine%d reading count:%d\n", i, cnt)
	time.Sleep(time.Millisecond)
	mr.RUnlock()
	fmt.Printf("goroutine%d reader over\n", i)
}

func write(mr *sync.RWMutex, i int) {
	fmt.Printf("goroutine%d writer start\n", i)
	mr.Lock()
	cnt++
	fmt.Printf("goroutine%d writing count:%d\n", i, cnt)
	time.Sleep(time.Millisecond)
	mr.Unlock()
	fmt.Printf("goroutine%d reader over\n", i)
}
