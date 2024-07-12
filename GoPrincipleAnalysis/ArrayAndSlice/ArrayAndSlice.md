# 数组和切片
Go语言中数组是一个值，数组变量表示整个数组，而不是和C语言一样指向第一个数组元素的指针。

## 有了数组为什么还需要切片？
1. 动态扩容
2. 动态数据集合处理问题

## 切片扩容
1. 计算目标容量
- case1：如果新切片的长度 > 旧切片容量的两倍，则新切片容量就为新切片的长度
- case2：
  - i：如果旧切片的容量小于256，那么新切片的容量就是旧切片的容量的两倍
  - ii：反之需要用旧切片容量按照1.25倍的增速，直到 >= 新切片长度 （为了更平滑的过渡，每次扩大1.25倍，还会加上3/4 * 256）
2. 内存对齐
需要按照Go内存管理的级别去对齐内存，最终容量以这个为主。

## 切片通过函数，传递的是什么？
值传递。
```go
type slice struct {
	array unsafe.Pointer
	len int
	cap int
}
```
详见Example1

## 在函数里面改变切片，函数外的切片会被影响吗？
- 底层数组没有改变，内外切片共享一个底层数组。
- 底层数组发生改变（如append超过底层数组容量，底层数组扩容），由于值传递，内外切片不再共享一个底层数组。

## 截取切片
- 截取切片只是修改了底层数组的起始指向和末尾指向，长度为截取长度，容量为起始指向到底层数组的末尾大小
- 用截取切片初始化新切片也是值传递

## 删除元素
- 删除某个元素，是底层数组中的元素逐个向前复制覆盖这个元素，数组的长度会减去删除元素的个数，但是底层数组的容量并没有发生变化
- 切片[]访问方法的大小是数组的长度，而不是容量，访问超过切片长度会报panic

## 新增元素
- append操作
- 触发扩容，按扩容规则扩容，返回新切片
- 未触发扩容，切片后新增，长度增加

## 深度拷贝
- 切片的传递，底层数组是浅拷贝（操作原来切片会影响新切片）
- 深度拷贝方法：copy(s2, s1)，复制的元素数量为，min(len(s1), len(s2))