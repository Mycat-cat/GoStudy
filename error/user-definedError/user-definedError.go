package main

import "fmt"

type MyError struct {
	code int
	msg  string
}

func (m MyError) Error() string {
	return fmt.Sprintf("code:%d,msg:%v", m.code, m.msg)
}

func NewError(code int, msg string) error {
	return MyError{
		code: code,
		msg:  msg,
	}
}

func Code(err error) int {
	if e, ok := err.(MyError); ok {
		return e.code
	}
	return -1
}

func Message(err error) string {
	if e, ok := err.(MyError); ok {
		return e.msg
	}
	return ""
}

func main() {
	err := NewError(100, "test MyError")
	fmt.Printf("code is %d,msg is %s", Code(err), Message(err))
	fmt.Printf("code is %d,msg is %s", err.(MyError).code, err.(MyError).msg)

	//err1 := MyError{100, "test MyError"}
	//s := err1.Error()
	//fmt.Println(s)
	//fmt.Println(err1.Error())

	//err := NewError(100, "test MyError")
	//fmt.Println(err.Error())
}
