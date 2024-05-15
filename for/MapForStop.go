package main

import "fmt"

func main() {
	v := make(map[int]int)
	v[1] = 1
	fmt.Println(v[0])
	for i, _ := range v {
		//fmt.Println(i)
		v[i+1] = i + 1
	}
	fmt.Println(v)
}
