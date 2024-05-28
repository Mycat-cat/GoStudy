package main

import "fmt"

/*
我们在每个goroutine中，向管道里发送一条数据，这样我们在程序最后，通过for循环将管道里的数据
全部取出，直到数据全部取出完毕才能继续后面的逻辑，这样就可以实现等待各个goroutine执行完。

但是，这样使用channel显得并不优雅，其次，我们得知道具体循环的次数，来创建管道的大小，假设
次数非常的多，则需要申请同样数量大小的管道出来，对内存也是不小的开销。
*/
func main() {
	ch := make(chan struct{}, 10)
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Printf("num:%d\n", i)
			ch <- struct{}{}
		}(i)
	}

	for i := 0; i < 10; i++ {
		<-ch
	}
	fmt.Println("end")
}
