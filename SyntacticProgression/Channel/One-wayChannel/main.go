package main

import (
	"fmt"
	"time"
)

/*
创建一个channel ch，分别定义两个单向channel类型SChannel和RChannel
根据别名类型给ch定义两个别名send和rec，一个只用于发送，一个只用于读取
*/

type SChannel = chan<- int
type RChannel = <-chan int

func main() {
	var ch = make(chan int)
	go func() {
		var send SChannel = ch
		fmt.Println("send:100")
		send <- 100
	}()

	go func() {
		var rec RChannel = ch
		num := <-rec
		fmt.Printf("receive:%d", num)
	}()
	time.Sleep(2 * time.Second)
}
