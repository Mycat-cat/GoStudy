<!-- TOC -->
* [channel](#channel)
  * [channel是什么](#channel是什么)
  * [channel数据结构](#channel数据结构)
  * [channel操作](#channel操作)
    * [channel初始化](#channel初始化)
    * [channel写入](#channel写入)
    * [channel读取](#channel读取)
    * [channel关闭](#channel关闭)
  * [Selcet](#selcet)
<!-- TOC -->
# channel
## channel是什么
- 通信管道，被设计用于实现goroutine之间的通信
- 以通信的方式来共享内存，而不是通过共享内存来实现通信

## channel数据结构
```go
type hchan struct {
	qcount   uint               // Channel环形数组中元素的数量 
	dataqsiz uint               // Channel环形数组的容量
	buf      unsafe.Pointer     // 指向Channel环形数组的一个指针
	elemsize uint16             // 元素所占的字节数
	closed   uint32             // 是否关闭
	elemtype *_type             // 元素类型
	sendx    uint               // 下一次写的位置
	recvx    uint               // 下一次读的位置
	recvq    waitq              // 读等待队列
	sendq    waitq              // 写等待队列
	lock mutex                  // runtime.mutex，保证Channel并发安全
}
```
recvq和sendq的结构，waitq结构体：
```go
type waitq struct {
	first *sudog    // sudog队列的队头指针
	last  *sudog    // sudog队列的队尾指针
}
```
sudog结构体：（可以看作是对阻塞挂起的g的一个封装，然后用多个sudog来构成等待队列）
```go
type sudog struct {
	// The following fields are protected by the hchan.lock of the
	// channel this sudog is blocking on. shrinkstack depends on
	// this for sudogs involved in channel ops.

	g *g   // 绑定的goroutine

	next *sudog     //前后指针
	prev *sudog
	elem unsafe.Pointer // 存储元素的容器

	// The following fields are never accessed concurrently.
	// For channels, waitlink is only accessed by g.
	// For semaphores, all fields (including the ones above)
	// are only accessed when holding a semaRoot lock.

	acquiretime int64
	releasetime int64
	ticket      uint32

	// isSelect indicates g is participating in a select, so
	// g.selectDone must be CAS'd to win the wake-up race.
	isSelect bool       // 是不是select操作封装的sudog

	// success indicates whether communication over channel c
	// succeeded. It is true if the goroutine was awoken because a
	// value was delivered over channel c, and false if awoken
	// because c was closed.
	success bool        // 为true，表示这个sudog是因为channel通信唤醒的
                        // 否则为false，表示这个sudog是因为channel close唤醒的
	// waiters is a count of semaRoot waiting list other than head of list,
	// clamped to a uint16 to fit in unused space.
	// Only meaningful at the head of the list.
	// (If we wanted to be overly clever, we could store a high 16 bits
	// in the second entry in the list.)
	waiters uint16

	parent   *sudog // semaRoot binary tree
	waitlink *sudog // g.waiting list or semaRoot
	waittail *sudog // semaRoot
	c        *hchan // 绑定的channel
}
```
**重点关注elem字段，elem作为收发数据的容器，当向channel发送数据时，elem代表将要写进channel的元素地址，当向channel读取数据时，elem代表要从channel中读取的元素地址**

## channel操作
### channel初始化
- 通过make函数来初始化一个channel
```go
func makechan(t *chantype, size int) *hchan {
    elem := t.Elem
    
    // compiler checks this but be safe.
    if elem.Size_ >= 1<<16 {
        throw("makechan: invalid channel element type")
    }
    if hchanSize%maxAlign != 0 || elem.Align_ > maxAlign {
        throw("makechan: bad alignment")
    }
    
    mem, overflow := math.MulUintptr(elem.Size_, uintptr(size))
    if overflow || mem > maxAlloc-hchanSize || size < 0 {
        panic(plainError("makechan: size out of range"))
    }
    
    // Hchan does not contain pointers interesting for GC when elements stored in buf do not contain pointers.
    // buf points into the same allocation, elemtype is persistent.
    // SudoG's are referenced from their owning thread so they can't be collected.
    // TODO(dvyukov,rlh): Rethink when collector can move allocated objects.
    var c *hchan
    switch {
    case mem == 0:
        // Queue or element size is zero.
        c = (*hchan)(mallocgc(hchanSize, nil, true))
        // Race detector uses this location for synchronization.
        c.buf = c.raceaddr()
    case elem.PtrBytes == 0:
        // Elements do not contain pointers.
        // Allocate hchan and buf in one call.
        c = (*hchan)(mallocgc(hchanSize+mem, nil, true))
        c.buf = add(unsafe.Pointer(c), hchanSize)
    default:
        // Elements contain pointers.
        c = new(hchan)
        c.buf = mallocgc(mem, elem, true)
    }

    c.elemsize = uint16(elem.Size_)
    c.elemtype = elem
    c.dataqsiz = uint(size)
    lockInit(&c.lock, lockRankHchan)

    if debugChan {
        print("makechan: chan=", c, "; elemsize=", elem.Size_, "; dataqsiz=", size, "\n")
    }
	return c
}
```
makechan函数有两个参数，第一个参数代表要创建的channel的元素类型，而第二个参数代表通道环形缓冲的容量大小

channel开辟内存分三种情况：
1. channel无缓冲 or 元素大小为0：只需要分配hchan本身结构体大小的内存
2. 有缓冲区buf，但元素不包含指针：hchan和buf一起分配
3. 有缓冲区buf，且元素包含指针类型：hchan和buf分开分配

### channel写入
```go
ch := make(chan int)
ch <- 1     // 往管道中写入1
```
写入的底层调用了runtime.chansend函数。往channel发送数据会出现三种情况：
1. channel中有读等待的goroutine
- 先拿锁
- 从recvq（读等待队列）里面弹出队列头部的sudog，进入send流程
- 将要写入的数据拷贝到这个sudog对应的elem数据容器上
- 释放锁
- 唤醒sudog绑定的goroutine（即将goroutine重新放入gmp模型中，等待调度）
2. channel中没有读等待的goroutine，并且环形缓冲数组里面有剩余空间
- 先拿锁
- 将数据写入到sendx指向的位置中
- sendx++，qcount++
- 释放锁
3. channel中没有读等待goroutine，并且无剩余空间存放数据
- 锁保护
- 获取一个sudog结构，绑定对应的channel，goroutine，还有ep指针
- 将sudog放入channel的写等待队列（sendq）
- runtime.gopark（挂起当前goroutine，可以看作是解绑当前G和M，然后开启下一轮调度）
4. 特殊情况1：写入的channel为nil
- 当channel为nil，对channel进行写操作，会导致当前goroutine永久性挂起，如果当前goroutine是main goroutine的话，还会导致整个程序退出
5. 特殊情况2：channel已经关闭，还想进行写操作
- 当channel已经关闭，再向channel写数据，会出现panic

### channel读取
```go
ch := make(chan, int)
v := <-ch       // 直接读取
v, ok <- ch     // ok判断读取的v是否有效
```
底层实际调用了runtime.chanrecv函数，往channel读取数据会出现三种情况
1. channel中有写等待goroutine
- 先拿锁
- 从sendq（写等待队列）里面弹出队列头部的sudog，进入recv流程
- 如果channel无缓冲区，直接读取sudog里面的数据，并唤醒sudog对应的goroutine
- 如果channel有缓冲区，读取环形缓冲区头部元素，并将sudog中的元素写入到缓冲区，唤醒sudog对应的goroutine
- 释放锁
2. channel中没有写等待goroutine，并且环形缓冲数组里面有剩余元素
- 先拿锁
- 读取recvx指向的数据
- recvx++，qcount--
- 释放锁
3. channel中没有写等待goroutine，并且环形缓冲数组里面无剩余元素
- 锁保护
- 获取一个sudog结构，绑定对应的channel，goroutine，还有ep指针
- 将sudog放入channel的读等待队列（recvq）
- runtime.gopark（挂起当前的goroutine，可以看作是解绑当前G和M，然后开启下一轮调度）
4. 特殊情况1：读取的channel为nil
- 当channel为nil，对channel进行读操作，会导致当前goroutine永久性挂起，如果当前goroutine是main goroutine的话，还会导致整个程序退出
5. 特殊情况2：channel已经关闭，并且buf里面没有元素
- channel已经关闭，并且没有剩余元素，还想读取channel会得到对应类型的零值

### channel关闭
```go
ch := make(chan int)
close(ch)
```
底层调用了runtime.closechan函数。具体流程：
- 如果对一个nil的channel执行close操作，会发生panic
- 加锁
- 如果重复关闭channel，也会panic
- 关闭channel（c.closed置为1）
- 将sendq和recvq里面所有等待者加入到glist中
- 唤醒glist中所有等待者（唤醒sudog对应的goroutine）

## Selcet
多路select，指一个goroutine可以服务多个channel的读或写操作

select分为两种：
1. 非阻塞型select（包含default分支的）
```go
package main

func main()  {
    ch := make(chan int)
	select {
	case <-ch1:
		
	case ch2 <- 1:

	default:
        
    }
}
```
2. 阻塞型select（不包含default分支的）
```go
package main

func main() {
    ch := make(chan int)
    select {
    case <-ch:
		
	case  ch <- 1:
    }
}
```
select核心原理：按照随机的顺序执行case，直到某个case完成操作，如果所有case都没有完成操作，则看有没有default分支，如果有default分支，则直接走default，防止阻塞。

如果没有的话，需要将当前goroutine加入到所有case对应channel的等待队列中，并挂起当前goroutine，等待唤醒。

如果当前goroutine被某一个case上的channel操作唤醒后，还需要将当前goroutine从所有case对应channel的等待队列中剔除。