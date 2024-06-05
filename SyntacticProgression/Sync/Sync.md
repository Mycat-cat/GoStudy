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

## sync/Atomic
### atomic和mutex的区别

操作方式：

- mutex：用于保护一段执行逻辑
- atomic：对变量进行操作

底层实现：

- mutex：操作系统调度器实现
- atomic：底层硬件指令支持，保证在cpu上执行不中断，atomic的性能也能随cpu的个数增加线性提升
```go
func AddT(addr *T, delta T)(new T)
func StoreT(addr *T, val T)
func LoadT(addr *T)(val T)
func SwapT(addr *T, new T)(old T)
func CompareAndSwapT(addr *T, old, new T)(swapped bool)
```

### atomic.value
atomic既可针对基本数据类型做原子操作，也可对多个变量进行同步保护，如struct复合类型。
提供:
- Load：从value读出数据
```go
func (v *Value) Load() (val any)
```
- Store：向value写入数据
```go
func (v *Value) Store (val any)
```
- Swap：用new交换value中存储的数据，返回value原来存储的旧数据
```go
func (v *Value) Swap (new any) (old any)
```
- CompareAndSwap：比较value中存储的数据和old是否相同，相同的话，将value中的数据替换为new
```go
func (v *Value) CompareAndSwap (old, new any) (swapped bool)
```

## sync.pool
sync.Pool是在sync包下的一个内存池组件，用来实现对象的复用，避免重复创建相同的对象，造成频
繁的内存分配和gc，以达到提升程序性能的目的。虽然池子中的对象可以被复用，但是是sync.Pool并
不会永久保存这个对象，池子中的对象会在一定时间后被gc回收，这个时间是随机的。所以，用
sync.Pool来持久化存储对象是不可取的。
另外，sync.Pool本身是并发安全的，支持多个goroutine并发的往sync.Poo存取数据
### sync.pool使用方法
- New()：sync.Pool的构造函数，用于指定sync.Pool中缓存的数据类
  型，当调用Get方法从对象池中获取对象的时候，对象池中如
  果没有，会调用New方法创建一个新的对象
- Get()：从对象池取对象
- Put()：往对象池放对象

**注意：取出放入之前记得Reset，不然初始对象无法复用**
### sync.pool使用场景
1. sync.pool主要是通过对象复用来降低gc带来的性能损耗，所以在高并发场景下，由于每个
   goroutine都可能过于频繁的创建一些大对象，造成gc压力很大。所以在高并发业务场景下出现
   GC 问题时，可以使用 sync.Pool 减少 GC 负担
2. sync.pool不适合存储带状态的对象，比如socket 连接、数据库连接等，因为里面的对象随时可能会被gc回收释放掉
3. 不适合需要控制缓存对象个数的场景，因为Poo池里面的对象个数是随机变化的，因为池子里的
   对象是会被gc的，且释放时机是随机的