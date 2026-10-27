package main

import "fmt"

func main() {

	//var challenge string
	//= 是单纯赋值，:= 是声明并赋值
	//challenge = "#90DaysOfDevOps"
	challenge := "#90DaysOfDevOps"
	//常量的声明不能省略 const 关键字，并且必须显式地初始化
	const daystotal = 90

	fmt.Printf("Welcome to %v\n", challenge)
	fmt.Printf("This is a %v challenge\n", daystotal)

	var Name string
	var DaysComplete int
	// ask user for their twitter handle

	Name = "王亚鹏"
	DaysComplete = 12
	fmt.Printf("%v has completed %v days of the challenge\n", Name, DaysComplete)
	fmt.Println("Great work")
}
