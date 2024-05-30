# Sync
Channel中我们提到，Go语言并发编程中，我们倡导通信共享内存，不使用共享内存通信，
Goroutine采用Channel来协作。但Go也提供了对共享内存并发安全机制的支持，这些内容都
在Sync包下。

## sync.WaitGroup
Goroutine示例中，我们都是利用time.Sleep()的方法让主goroutine等待以便子goroutine能够执行完打印结果。
然而你并不知道所有的子goroutine什么时候执行完，不知道确切的要等多久。

如何处理？
### Channel
参见SyncExample1。
### sync.WaitGroup
Go语言中可以使用sync.WaitGroup来实现并发任务的同步以及协程任务等待。
```go
//sync.WaitGroup是一个对象,里面维护着一个计数器，并且通过三个方法来配合使用
(wg * WaitGroup) Add(delta int)  // 计数器加delta
(wg * WaitGroup) Done()          // 计算器减1  
(wg * WaitGroup) Wait()          // 会阻塞代码的运行，直至计算器减为0
```
参见SyncExample2。
## sync.Once
在我们写项目的时候，程序中有很多的逻辑只需要执行一次，最典型的就是项目工程里配置文件的加载，我们只需要加载一次即可，让配置保存在内存中，下次使用的时候直接使用内存中的配置数据即可。这里就要用到sync.Once。

sync.Once可以在代码的任意位置初始化和调用，并且线程安全。sync.Once最大的作用就是延迟初始
化，对于一个sync.Once变量我们并不会在程序启动的时候初始化，而是在第一次用的它的时候才会初
始化，并且只初始化这一次，初始化之后驻留在内存里，这就非常适合我们之前提到的配置文件加载
场景，设想一下，如果是在程序刚开始就加载配置，若迟迟未被使用，则既浪费了内存，又延长了程
序加载时间，而sync.Once就刚好解决了这个问题。

参见SyncExample3

### 与init()的区别
init()方法是在其所在的package首次加载时执行
sync.Once可以在代码的任意位置初始化和调用，是在第一次用它的时候才会初始化

## sync.Lock
Go语言中有两种方式来控制并发安全，锁和原子操作

参见SyncExample4，其中涉及到并发问题，同一时间有多个goroutine在对num做+1操作，但是存在后一个不是在前一个+1完成的基础上运行的，所以就导致了num被覆盖，最终不等于n的值。

为了避免上述的并发安全问题，一般采用以下两种方式处理。

### 锁
#### 互斥锁Mutex
互斥锁是一种最常用的控制并发安全的方式，它在同一时间只允许一个goroutine对共享资源进行访问。
```go
var lock sync.Mutex
```
互斥锁有两个方法：
```go
func (m *Mutex) Lock()  // 加锁
func (m *Mutex) UnLock()  // 解锁
```
一个互斥锁只能同时被一个goroutine锁定，其它goroutine将阻塞直到互斥锁被解锁才能加锁成功

sync.Mutex在使用时注意：
**对一个未锁定的互斥锁解锁将会产生运行时错误**
参见SyncExample5

#### 读写锁RWMutex
将读操作与写操作分开，可以分别对读和写进行加锁，一般用在大量读操作，少量写操作的情况
方法如下：
```go
func(rw *RWMutex) Lock()    // 对写锁加锁
func(rw *RWMutex) Unlock()  // 对写锁解锁  
func(rw *RWMutex) RLock()   // 对读锁加锁
func(rw *RWMutex) RUnlock() // 读读锁解锁
```
读写锁的使用一般遵循以下几个法则：
1. 同时只能有一个goroutine能够获得写锁定
2. 同时可以有任意多个goroutine获得读锁定
3. 同时只能存在写锁定或读锁定（读和写互斥）

读写锁示例：
参见SyncExample6

### 死锁
死锁是一种状态，当两个或两个以上的goroutine在执行过程中，因抢夺共享资源处在互相等待的状态。如果没有外部干涉将会一直处于这种阻塞状态，我们称这时的系统发生了死锁。

#### 死锁场景
##### Lock/UnLock不成对
最常见场景就是对锁进行拷贝使用。
参见SyncExample7

在使用锁的时候，我们应当尽量避免锁拷贝，并且保证Lock()和Unlock()成对出现，
没有成对出现容易会出现死锁的情况，或者是Unlock一个未加锁的Mutex而导致panic。

```go
mu.Lock()
defer mu.Unlock()
```

##### 循环等待
循环等待造成死锁：如A等B，B等C，C等A，循环等待

参见SyncExample8

## sync.Map
map不能同时被多个goroutine读写
解决方法：
1. 对map加锁
2. sync.Map(Go 1.9引入)

sync.Map无需初始化即可使用，内置操作方法。

**sync.Map没有提供获取map数量的方法，需要我们对其进行遍历计算，为了保证并发安全存在性能损失，在非并发情况下，map性能优于sync.Map**

