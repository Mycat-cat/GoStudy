package main

import "fmt"

/*
函数的return并非原子操作,return的过程可以被分解为以下三步：
1. 设置返回值
2. 执行defer语句
3. 将结果返回

本例中,第一步将res的值设置为num,此时还未执行defer,num的值为1,所以res被设置为1,然后再执行defer语句将result+1,最终将result返回,打印出2
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
