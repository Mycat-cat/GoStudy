package main

import (
	"fmt"
	"time"
)

func add(num *int) {
	*num = *num + 1
}

/*
不加锁，循环100次结果还正确是因为100太小了对计算机来说非常简单
当循环100000w就可以发现结果不对了
*/
func main() {
	var num int
	for i := 0; i < 100000; i++ {
		go add(&num)
	}
	time.Sleep(2)
	fmt.Println("num的值：", num)
}
