package main

import "fmt"

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

	var phone Phone
	//phone = new(Apple)
	//phone = &Apple{}
	//phone.(*Apple).PhoneName = "Apple"

	// 断言成功但是无法修改Apple的PhoneName,除非传递指针类型
	phone = Apple{"Apple"}
	val, ok := phone.(Apple)
	fmt.Println(val, ok)

	//phone.Call()
	//phone.SendMessage()
}
