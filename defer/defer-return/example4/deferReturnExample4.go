package main

import "fmt"

/*
函数的return并非原子操作,return的过程可以被分解为以下三步：
1. 设置返回值
2. 执行defer语句
3. 将结果返回

本例中,第一步将res的值设置为num,此时还未执行defer,num的值为1,所以res被设置为1,然后再执行defer语句将res+1,最终将res返回,打印出2

还存在一个疑问？为什么匿名函数func()能够捕获到res？匿名函数不是可以捕捉外部作用域中的变量,res也是外部作用域中的变量吗？
回答：是的,defRun函数定义了一个命名返回值res int,在Go中,命名返回值在函数开始执行时就被初始化了。此例中,res的初始值为0
*/
func deferRun() (res int) {
	num := 1
	defer func() {
		res++
	}()
	return num
}

func main() {
	res := deferRun()
	fmt.Println(res)
}
