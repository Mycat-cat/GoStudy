package main

import "fmt"

func deferRun() {
	var num = 1
	numPtr := &num
	defer fmt.Println("defer1:", numPtr, num, *numPtr)
	defer func() {
		fmt.Println("defer2", numPtr, num, *numPtr)
	}()
	defer func(num int, numPtr *int) {
		fmt.Println("defer3", numPtr, num, *numPtr)
	}(num, numPtr)

	num = 2
	fmt.Println(numPtr, num, *numPtr)
}

func main() {
	deferRun()
}
