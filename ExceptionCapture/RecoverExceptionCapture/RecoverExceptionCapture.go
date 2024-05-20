package main

import "fmt"

/*
这里有了recover之后,程序不会在panic处中断,在执行完panic之后,会接下来执行defer recover函数,但是当前函数panic后面的代码
不会被执行,但是调用该函数的代码会接着执行.
如果我们在main函数中未加入defer fun(){...},当我们的程序运行到panic时会崩溃掉,为了保障程序健壮的运行,而不是因为一些panic挂掉
所以当发生panic的时候我们要让程序能够继续运行,并且获取到发生panic的具体错误
*/

func main() {
	defer func() {
		if error := recover(); error != nil {
			fmt.Println("出现了panic,使用recover获取信息：", error)
		}
	}()
	fmt.Println("11111111")
	panic("出现panic")
	fmt.Println(22222222)
}
