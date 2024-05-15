package main

import "fmt"

const m = iota
const n = iota
const q = 1
const (
	x = iota
	y
	z
)

const (
	x1 = 1
	y1
	z1
)

func main() {
	const (
		a = iota
		b
		c
	)
	const (
		d = 1
		e
		f
	)

	const (
		g = iota
		h
		i
	)
	fmt.Println(a, b, c)
}
