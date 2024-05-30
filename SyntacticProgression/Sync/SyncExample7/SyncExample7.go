package main

import (
	"fmt"
	"sync"
)

/*
会报死锁：
这里mu sync.Mutex当作参数传入到函数copyMutex,
锁进行了拷贝，不是原来的锁变量了，那么一把新的锁，在执行mu.Lock()的时候应该没问题。这就是要注意的地方，
如果将带有锁结构的变量赋值给其他变量，锁的状态会复制。
所以多锁复制后的新的锁拥有了原来的锁状态，那么在copyMutex函数内执行mu.Lock()的时候会一直阻塞，因为外层的main
函数已经Lock()了一次，但是并没有机会Unlock()，导致内层函数会一直等待Lock()，而外层函数一直
等待Unlock()，这样就造成了死锁。
*/

func main() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()
	copyMutex(mu)
}
func copyMutex(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("ok")
}
