package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var A uint = 60
	var B uint = 13
	var C uint
	var D uint
	var E int
	var F int32
	var G bool
	var H int64
	var (
		I complex64
		J complex128
		K string
		L float32
		M float64
	)
	C = A ^ B
	D = ^A
	fmt.Println(C, D)
	fmt.Println(unsafe.Sizeof(A), unsafe.Sizeof(E), unsafe.Sizeof(F), unsafe.Sizeof(G), unsafe.Sizeof(H))
	fmt.Println(unsafe.Sizeof(I), unsafe.Sizeof(J), unsafe.Sizeof(K))
	fmt.Println(unsafe.Sizeof(L), unsafe.Sizeof(M))

	fmt.Printf("%T", A)
}
