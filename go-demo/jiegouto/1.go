package main

import "fmt"

// 定义接口 Speaker：约定必须有 Speak() 方法
type Speaker interface {
	Speak() string
}

type Person struct {
	Name string
}

func (p Person) Speak() string {
	return "哈喽我是 " + p.Name
}

type Dog struct {
}

func (d Dog) Speak() string {
	return "汪汪汪！"
}

// main 函数是程序的入口点
func main() {

	var s Speaker
	s = Person{Name: "王亚鹏"}
	fmt.Println(s.Speak())

	s = Dog{}
	fmt.Println(s.Speak())
}
