package main

import (
	"fmt"
)

// 主函数入口
func main() {

	// 定义常量，总天数为90天
	const DaysTotal int = 90
	// 定义挑战主题字符串
	challenge := "#90DaysOfDevOps"

	// 打印欢迎信息和挑战详情
	fmt.Printf("Welcome to the %v challenge.\nThis challenge consists of %v days\n", challenge, DaysTotal)

	// 声明变量用于存储用户输入的Twitter用户名
	var TwitterName string
	// 声明变量用于存储用户完成的挑战天数
	var DaysCompleted uint

	// 提示用户输入Twitter用户名
	// asking for user input
	fmt.Println("Enter Your Twitter Handle: ")
	// 读取用户输入的Twitter用户名
	fmt.Scanln(&TwitterName)

	// 提示用户输入已完成的挑战天数
	fmt.Println("How many days have you completed?: ")
	// 读取用户输入的已完成天数
	fmt.Scanln(&DaysCompleted)

	// 打印感谢信息，包含用户名和已完成天数
	fmt.Printf("Thank you %v for taking part and completing %v days.\n", TwitterName, DaysCompleted)
	// 打印祝福语
	fmt.Println("Good luck")
}
