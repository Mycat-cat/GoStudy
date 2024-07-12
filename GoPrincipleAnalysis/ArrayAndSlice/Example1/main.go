package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	s := make([]int, 5, 10)
	PrintSliceStruct(&s)
	test(s)
}
func test(s []int) {
	PrintSliceStruct(&s)
}

func PrintSliceStruct(s *[]int) {
	ss := (*reflect.SliceHeader)(unsafe.Pointer(s))
	fmt.Printf("Slice struct: %+v, slice is %v\n", ss, s)
}
