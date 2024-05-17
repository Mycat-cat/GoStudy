package main

import "fmt"

/*
函数的return并非原子操作,return的过程可以被分解为以下三步：
1. 设置返回值
2. 执行defer语句
3. 将结果返回

本例中,与example4的区别在于返回值匿名,第一步将返回值设置为1,此时还未执行defer,num的值为0,然后再执行defer语句将num+1,最终将返回值返回,打印出1
*/
func deferRun() int {
	var num int
	defer func() {
		num++
	}()

	return 1
}

func main() {
	res := deferRun()
	fmt.Println(res)
}
