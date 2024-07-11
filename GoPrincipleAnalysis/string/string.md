<!-- TOC -->
* [string](#string)
  * [string与[]byte的转化原理](#string与byte的转化原理)
  * [字符串声明](#字符串声明)
  * [字符串拼接](#字符串拼接)
<!-- TOC -->
# string
- 字符串是所有8bit字节的集合，但不一定是UTF-8编码的文本
- 字符串可以是empty，但不能是nil，empty字符串是一个没有任何字符的空串""
- 字符串不可以被修改，字符串类型不可变（赋值为重新开辟存储空间，原来的空间并不会立即释放，需要GC）

```go
type stringStruct struct {
	str unsafe.Pointer
	len int
}
```
- str：字符串首地址
- len：字符串长度（实际字节数，对于非单字节编码字符，其结果可能多于字符个数）

字符串上的写操作包括拼接、追加通过拷贝实现

字符串修改并不等于重新赋值。
```go
str := "Hello"
str := "Golang"  // 重新赋值
str[0] = "I"     // 修改，不允许
```

可以将字符串转化为字节数组，然后通过下标修改字节数组，再转化回字符串。
详见stringExample1（最终得到的只是ss字符串的一个拷贝，源字符串没有变化）；

## string与[]byte的转化原理
- 发生一次内存拷贝，并申请一块新的切片内存空间

什么情况不发生内存拷贝？
答：转化为的字符串被用于临时场景，string的指针（string.str）指向切片的内存

举例：
1. 字符串比较：string(ss) == "Hello"
2. 字符串拼接："Hello" + string(ss) + "World"
3. 用作查找：比如map的 key, val := map[string(ss)]

## 字符串声明
```go
str1 = "Hello World"
str2 = `Hello
```
双引号字符串只能用于单行字符串的初始化，特殊字符需要转义。

反引号声明的字符串没有限制，字符内容即字符串原始内容，一般用反引号表示比较复杂的字符串，如json串。

## 字符串拼接
性能对比：strings.builder ≈ strings.join > bytes.buffer > append > "+" > fmt.sprintf

少量字符串拼接：直接使用+操作符是最方便也算是性能最高的，就无需使用strings.builder