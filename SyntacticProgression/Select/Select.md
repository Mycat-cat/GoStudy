<!-- TOC -->
* [Select](#select)
  * [Select用法](#select用法)
    * [空select永久阻塞](#空select永久阻塞)
    * [没有default且case无法执行的select永久阻塞](#没有default且case无法执行的select永久阻塞)
    * [有单一case和default的select](#有单一case和default的select)
    * [有多个case和default的select](#有多个case和default的select)
<!-- TOC -->
# Select
Go语言层面提供的一种多路复用机制，用于检测当前goroutine连接的多个channel是否有数据准备完毕，可用于读或写

Go语言的select：起一个goroutine监听多个Channel的读写事件，提高从多个Channel获取信息的效率，相当于是单线程处理多个IO事件

## Select用法
```go
select {
    case <- channel1:
       do ...
    case channel2 <- 1:
       do ...
    default:
       do ...
}
```
select的各个case的表达式必须都是channel的读写操作。
select通过多个case语句监听多个channel的读写操作是否准备好可以执行，其中任何一个case看可以执行了则选择该case语句执行，如果没有可以执行的case，则执行default语句，如果没有default，则当前goroutine阻塞，如果有多个case都可以执行，则随机选择一个执行。
### 空select永久阻塞
```go
func main() {
	select {
	}
}
```
### 没有default且case无法执行的select永久阻塞

```go
package main

import
import "fmt"
{
"fmt"
}

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)
	select {
	case <-ch1:
            fmt.Printf("received from ch1")
	case num := <-ch2:
            fmt.Printf("num is: %d", num)
    }
}
```
程序中select从两个channel，ch1和ch2中读取数据，但是两个channel都没有数据，且没有goroutine往里面写数据，所以不可能读到数据，这两个case永远无法执行到，select也没有default，所以会出现永久阻塞，报死锁。

### 有单一case和default的select
```go
package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1)
	select {
	case <-ch:
            fmt.Println("received from ch")
	default:
            fmt.Println("default!!!")
    }
}
```
执行到select语句的时候，由于ch中没有数据，且没有goroutine往channel中写数据，所以不可能执行到，就会执行default语句，打印出default!!!

### 有多个case和default的select
当多个case都准备好了的时候，随机选择一个执行。

select语句不在循环内，只会执行一次，所以当多个case都准备好了的时候，select选择的随机性会导致随机输出一个case
