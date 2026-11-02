package main

import "fmt"

//1结构体定义
/*type Student struct {
	Name  string
	Score int
}

func main() {
	student := Student{Name: "小林", Score: 80}
	fmt.Println(student.Name, student.Score)
	fmt.Println(student)
}*/

//2通过结构体指针修改原数据
/*type Student struct{ Score int }

func setScore(student *Student) {
	student.Score = 85 //go会自动解引用，所以写成(*student).Score = 85
}
func main() {
	s := Student{Score: 80}
	setScore(&s) //要想在函数内部改变外部结构体的状态，必须传递它的内存地址（指针）。
	fmt.Println(s.Score)
}*/

//3 俩份值与一个指针   克隆体与本体
type Student struct{ Score int }

func main() {
	first := Student{Score: 80}
	second := first   //复制
	pointer := &first //取地址
	second.Score = 90
	pointer.Score = 85
	fmt.Println(first.Score, second.Score, pointer.Score)
}
