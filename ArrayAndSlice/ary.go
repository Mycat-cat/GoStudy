package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var strAry = [10]string{"aa", "bb", "cc", "dd", "ee", "ff"}
	fmt.Printf("%v\n", strAry)
	fmt.Printf("%+v\n", strAry)
	fmt.Println(strAry)
	fmt.Println(len(strAry))

	var sliceAry = make([]string, 0)
	sliceAry = strAry[1:3]

	var dic = map[string]int{
		"apple":      1,
		"watermelon": 2,
	}

	fmt.Println(sliceAry)
	fmt.Println(dic)
	fmt.Printf("%v\n", dic)
	fmt.Printf("%+v\n", dic)

	p := Person{Name: "Alice", Age: 30}
	fmt.Printf("%v\n", p)
	fmt.Printf("%+v\n", p)
}
