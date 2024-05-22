package main

import (
	"fmt"
)

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go func() {
		sum(s[:len(s)/2], c)
	}()
	//go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)
}

/*
输出 -5, 17, 12
goroutine 的顺序是不定的，这里我做了一个小测试，让整个程序循环了 10W 次，统计 x=17
和-5 的次数。
结果表明，x=17与x=-5 的概率
大概是 1:8。
*/
