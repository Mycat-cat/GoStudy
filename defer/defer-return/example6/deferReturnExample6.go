package main

import "fmt"

/*
函数的return并非原子操作,return的过程可以被分解为以下三步：
1. 设置返回值
2. 执行defer语句
3. 将结果返回

本例中,返回值匿名,第一步将返回值设置为1,此时还未执行defer,num的值为1,然后再执行defer语句将num+1,最终将返回值返回,打印出1
*/
func deferRun() int {
	num := 1
	defer func() {
		num++
	}()

	return num
}

func main() {
	res := deferRun()
	fmt.Println(res)
}
