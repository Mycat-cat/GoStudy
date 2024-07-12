package main

import (
	"bytes"
	"fmt"
	"reflect"
	"strings"
	"unsafe"
)

func main() {
	// 构建strings.Builder
	a := []byte{1, 2, 3}
	b := strings.Builder{}
	b.Write(a)

	// 构建bytes.Buffer
	b2 := bytes.NewBuffer(a)

	str1 := b.String()
	str2 := b.String()

	// 通过反射获取str1,str2的底层描述（比如它的数组指针）：数组指针一致（复用内存）
	String2Bytes(str1)
	String2Bytes(str2)

	str3 := b2.String()
	str4 := b2.String()

	// 通过反射获取str3,str4的底层描述（比如它的数组指针）：数组指针不一致（没有复用）
	String2Bytes(str3)
	String2Bytes(str4)
}

func String2Bytes(s string) {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{
		Data: sh.Data,
		Len:  sh.Len,
		Cap:  sh.Len,
	}
	fmt.Println(bh.Data)
}
