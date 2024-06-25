<!-- TOC -->
* [定时器](#定时器)
  * [Timer](#timer)
    * [创建Timer](#创建timer)
    * [停止Timer](#停止timer)
    * [重置Timer](#重置timer)
    * [time.AfterFunc](#timeafterfunc)
    * [time.After](#timeafter)
* [Ticker](#ticker)
<!-- TOC -->
# 定时器
定时任务
## Timer
一次性时间定时器，即在未来某个时刻，触发的事件只会执行一次。
```go
type Timer struct {
	C <-chan Time
	r runtimeTimer
}
```
Time类型的管道C，主要用于事件通知，在未到达设定时间，管道内没有数据写入，阻塞，到达设定时间，会像管道写入一个系统时间，触发事件
### 创建Timer
```go
func NewTimer(d Duration) *Timer
```
### 停止Timer
```go
func (t *Timer) Stop() bool
```
返回值：
- true：超时时间内停止Timer
- false：超时时间外停止了Timer

### 重置Timer
```go
func (t *Timer) Reset(d Duration) bool
```
对于已经过期或者是已经停止的timer，可以通过该方式激活使其继续生效

### time.AfterFunc
```go
func AfterFunc(d Duration, f func()) *Timer
```
- d：超时时间d
- f：具体的函数
在创建出timer之后，在当前goroutine，等待一段时间d之后，执行f

### time.After
```go
func After(d Duration) <-chan Time {
	return NewTimer(d).C
}
```
After函数会返回timer里的管道，并且这个管道会在经过时段d之后写入数据，调用这个函数，就相当于实现了定时器。

# Ticker
```go
func NewTicker(d Duration) *Ticker
```
```go
type Ticker struct {
	C <-chan Time
	r runtimeTimer
}
```
Ticker字段与Timer一致，包含一个通道字段，每隔时间段d就向该通道发送当时的时间，根据管道消息来触发事件。Ticker定义完成，从当前时间开始计时，每隔固定时间都会触发，只有关闭Ticker对象才会停止。

