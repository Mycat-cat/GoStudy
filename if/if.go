package main

import "fmt"

func main() {
	var dic = map[string]int{
		"apple":      1,
		"watermelon": 2,
	}
	num, ok := dic["watermelon"]
	fmt.Println(num, ok)

	if num, ok := dic["orange"]; ok {
		fmt.Printf("orange num %d\n", num)
	}

	if num, ok := dic["apple"]; ok {
		fmt.Printf("apple num %d\n", num)
	}
}
