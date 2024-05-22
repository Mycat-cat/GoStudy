package main

func main() {
	//ch := make(chan bool, 1)
	//ch <- true
	//go func() {
	//	ch <- true
	//}()

	ch := make(chan bool, 1)
	ch <- true
	ch <- true
}

/*
上面的代码是子协程死锁，主协程正常结束，子协程跟随结束
下面的代码是主协程死锁
*/
