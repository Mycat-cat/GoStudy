package main

//import "fmt"
//
//func main() {
//	var a int = 5
//	res := []*int{&a}
//	fmt.Println(*res[0])
//
//	var b *int = &a
//	fmt.Println(*b)
//
//	var c int = 6
//	// press := []*int{},空切片,建议更改为下面的0切片声明形式
//	var press []*int
//	press = append(press, &a, &c)
//	fmt.Println(*press[0], *press[1])
//
//	// Go 1.22版本之后是可以取到arr元素的地址，1.22版本之前输出的结果应该是2,2
//	arr := [2]int{1, 2}
//	res = []*int{}
//	for _, v := range arr {
//		fmt.Println(&v)
//		res = append(res, &v)
//	}
//	fmt.Println(*res[0], *res[1])
//
//	// Go 1.22版本之前如何得到预期结果：1,2
//	// 方式1
//	res = []*int{}
//	for _, v := range arr {
//		v1 := v
//		fmt.Println(&v1)
//		res = append(res, &v1)
//	}
//	fmt.Println(*res[0], *res[1])
//
//	// 方式2
//	res = []*int{}
//	for k := range arr {
//		res = append(res, &arr[k])
//	}
//	fmt.Println(*res[0], *res[1])
//}
