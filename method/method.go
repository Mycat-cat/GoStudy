package main

import "fmt"

type Student struct {
	ID    int
	Name  string
	Age   int
	Score int
}

func (st *Student) SetScore(score int) {
	st.Score = score
}

func (st *Student) GetScore() int {
	return st.Score
}

func main() {
	st := &Student{
		ID:    100,
		Name:  "ZhangSan",
		Age:   18,
		Score: 98,
	}
	fmt.Printf("设置前，学生st的分数是：%d\n", st.GetScore())
	st.SetScore(100)
	fmt.Printf("设置后，学生st的分数是：%d\n", st.GetScore())
}
