package main

import "fmt"

func main() {
	const (
		a = 1
		b
		c
		d = iota
		e = iota
		f = iota
	)
	fmt.Println(a, b, c, d, e, f)
}
