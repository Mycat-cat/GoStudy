package main

import (
	"fmt"
)

func main() {
	var ss string
	ss = "Hello"
	strByte := []byte(ss)
	strByte[1] = 65
	fmt.Println(string(strByte))
}
