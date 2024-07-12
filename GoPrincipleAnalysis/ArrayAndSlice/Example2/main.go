package main

import "fmt"

func main() {
	doAppend := func(s []int) {
		s = append(s, 1)
		printLengthAndCapacity(s)
	}
	s := make([]int, 8, 8)
	doAppend(s[:4])
	printLengthAndCapacity(s)
	doAppend(s)
	printLengthAndCapacity(s)
}
func printLengthAndCapacity(s []int) {
	fmt.Println(s)
	fmt.Printf("len=%d cap=%d\n", len(s), cap(s))
}
