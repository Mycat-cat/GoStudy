# Go语言error

代表函数执行过程中或者逻辑有出错

## error

error是一个普通的接口类型，不携带任何堆栈信息

```go
type error interface {
    Error() string
}
```

error.New()返回的是一个地址,函数实现

```go
func New(text string) error {
    return &errorString{text}
}
```

fmt.Errorf()实现

```go
func Errorf(format string, a ...interface{}) error {
    p := newPrinter()
    p.wrapErrs = true
    p.doPrintf(format, a)
    s := string(p.buf)
    var err error
    if p.wrappedErr == nil {
        err = errors.New(s)
    } else {
        err = &wrapError{s, p.wrappedErr}	
    }
    p.free()
    return err
}
```

## 自定义error对象
