package main

import "fmt"

func deferRun() {
	var a *int
	var b = 200
	defer fmt.Println(*a)
	a = &b
	b = 200
}

func main() {
	deferRun()
}
