package main

import "fmt"

func deferRun() {
	var num = 1
	numAddr := &num
	//defer func(intAddr *int) {
	//	fmt.Println(*intAddr)
	//}(numAddr)

	defer func() {
		fmt.Println(*numAddr)
	}()

	num = 2
	fmt.Println(*numAddr)
	return
}

func main() {
	deferRun()
}
