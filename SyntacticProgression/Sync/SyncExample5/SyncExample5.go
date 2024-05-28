package main

import (
	"fmt"
	"sync"
)

var (
	num int
	wg  = sync.WaitGroup{}
	mu  = sync.Mutex{}
)

func add() {
	mu.Lock()
	defer wg.Done()
	num += 1
	mu.Unlock()
}

func main() {
	var n = 10 * 10 * 10 * 10
	wg.Add(n)
	for i := 0; i < n; i++ {
		go add()
	}
	wg.Wait()
	fmt.Println(num == n) // 理论上num应该等于n,但是实际上不一致
}
