# Go语言defer

defer关键字，主要用在函数或者方法前面，作用是用于函数和方法的延迟调用。

# defer常用场景

## 资源的释放

通过defer延迟调用机制，处理资源回收问题，如网络连接、数据库连接以及文件句柄的资源释放。

## 配合recover一起处理panic

Go语言中用panic来抛出异常，用recover来捕获异常。当程序出现异常，我们需要知道发生了什么异常的时候，可以用defer recover来捕获该异常。

# defer与return 

defer函数的执行是在return的时候。return的时候，defer具体做了什么？又会带来什么结果？

1. defer定义的延迟函数的参数在defer语句出时就已经确定下来了
2. return不是原子级操作，执行过程是：设置返回值->执行defer语句->将结果返回