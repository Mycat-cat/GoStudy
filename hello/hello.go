package main

import (
	"GoStudy/myMath"
	"fmt"
)

var a *int
var b []int
var c map[string]int
var d chan int
var e func(string) int
var f error

// 因式分解关键字的写法一般用于声明全局变量
var (
	x int
	y int
)

var l, m int = 1, 2
var n, o = 123, "hello"

func main() {
	fmt.Println("Hello, World!")
	fmt.Println(myMath.Add(1, 1))
	fmt.Println(myMath.Sub(1, 1))

	fmt.Println(a == nil)
	fmt.Println(b == nil)
	fmt.Println(c == nil)
	fmt.Println(d == nil)
	fmt.Println(e == nil)
	fmt.Println(f == nil)

	var g int
	var h float64
	var i bool
	var j string
	fmt.Println(g, h, i, j)
	fmt.Printf("%v %v %v %q\n", g, h, i, j)

	var k = true
	fmt.Println(k)

	// intVal := 1
	var intVal int = 1
	f := "Ru noob" // var f string = "Ru noob"

	/*
		出现在 := 左侧的变量不能是已经声明过的，:=是使用变量的首选形式，但是它只能被用在函数体内。
	*/

	fmt.Println(intVal)
	fmt.Println(f)

	// 声明类型相同多个变量，非全局变量
	var vname1, vname2, vname3 string
	vname1, vname2, vname3 = "I", "am", "HelloWorld"
	fmt.Println(vname1, vname2, vname3)

	var vname4, vname5, vname6 = "i", "am", "helloworld"
	fmt.Println(vname4, vname5, vname6)

	vname7, vname8, vname9 := "helloworld", "is", "me"
	fmt.Println(vname7, vname8, vname9)

	fmt.Println(x, y)

	var p, q = 123, "hello"
	fmt.Println(n, o)
	fmt.Println(p, q)

	// 只能在函数体内部使用
	r, s := 123, "hello"
	fmt.Println(r, s)
}
