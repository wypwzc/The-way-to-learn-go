package main

import "fmt"

//结构体定义
/*type Student struct {
	Name  string
	Score int
}

func main() {
	student := Student{Name: "小林", Score: 80}
	fmt.Println(student.Name, student.Score)
	fmt.Println(student)
}*/

//通过结构体指针修改原数据
type Student struct{ Score int }

func setScore(student *Student) {
	student.Score = 85 //go会自动解引用，所以写成(*student).Score = 85
}
func main() {
	s := Student{Score: 80}
	setScore(&s) //要想在函数内部改变外部结构体的状态，必须传递它的内存地址（指针）。
	fmt.Println(s.Score)
}
