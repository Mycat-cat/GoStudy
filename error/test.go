package main

import (
	"errors"
	"fmt"
)

func getPositiveSelfAdd(num int) (int, error) {
	if num <= 0 {
		return -1, fmt.Errorf("num is not a positive number")
	}
	return num + 1, nil
}

func main() {
	num1, err1 := getPositiveSelfAdd(1)
	fmt.Printf("num is %d,err is %v\n", num1, err1)

	num2, err2 := getPositiveSelfAdd(-2)
	fmt.Printf("num is %d,err is %v\n", num2, err2)

	err3 := errors.New("hello")
	err4 := errors.New("hello")
	fmt.Println(err3 == err4)

	fmt.Println(err3.Error() == err4.Error())
}
