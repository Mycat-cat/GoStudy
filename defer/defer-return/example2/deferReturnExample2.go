package main

import "fmt"

/*
defer出现的时候,参数已经确定,但这里传递的是地址,地址没变，地址对应的内容被修改了，所以输出被改变了
*/
func deferRun() {
	var arr = [4]int{1, 2, 3, 4}
	defer printArr(&arr)

	arr[0] = 100
	return
}

func printArr(arr *[4]int) {
	for i := range arr {
		fmt.Println(arr[i])
	}
}

func main() {
	deferRun()
}
