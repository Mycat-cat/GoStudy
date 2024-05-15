# Go 语言接口

```go
type interfaceName interface {
    methodName1([parameter_list]) [return_type_list]
    methodName2([parameter_list]) [return_type_list]
    ...
}
```

Go 语言中的接口是一组方法的声明。当某一种类型实现了所有这些声明的方法，则称这种类型为该接口的一种实现。

```go
type Phone interface {
    Call()
    SendMessage()
}

type Apple struct {
    PhoneName string
}

func (a Apple) Call() {
    fmt.Printf("%s有打电话功能\n", a.PhoneName)
}

func (a Apple) SendMessage() {
    fmt.Printf("%s有发短信功能\n", a.PhoneName)
}

type HuaWei struct {
    PhoneName string
}

func (h HuaWei) Call() {
    fmt.Printf("%s有打电话功能\n", h.PhoneName)
}

func (h HuaWei) SendMessage() {
    fmt.Printf("%s有发短信功能\n", h.PhoneName)
}

func main() {
    a := Apple{"apple"}
    b := HuaWei{"huawei"}
    a.Call()
    a.SendMessage()
    b.Call()
    b.SendMessage()
    
    var phone Phone        // 声明一个接口类型phone
    phone = new(Apple)     // 注意这种创建方式，new函数参数是接口的实现
    phone.(*Apple).PhoneName = "Apple"    // 这里使用断言给phone的成员赋值 
    phone.Call()
    phone.SendMessage()
}
```

1. 实现多个接口

类型可以实现接口，一种类型可以实现多个接口。

2. 空接口

没有任何方法声明的接口称为空接口。

```go
interface{}
```

所有的类型都实现了空接口，因此空接口可以存储任意类型的数值。

3. 断言

类型断言（Type Assertion）接口操作，用来检查接口变量的值是否实现了某个接口或者是否是某个具体的类型

```go
value, ok := x.(T)    // x为接口类型，ok为bool类型
```

4. 接口作函数参数

接口作函数参数，在函数定义的时候，形参为接口类型，在函数调用的时候，实参为该接口的具体实现。

5. 接口嵌套
