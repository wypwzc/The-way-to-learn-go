package main

import "fmt"

// 定义结构体类型 User
type User struct {
	name   string
	age    int
	height float64
}

func main() {

	/*var challenge = "#90DaysOfDevOps"

	const daystotal = 90

	var dayscomplete = 11

	fmt.Println("Welcome to", challenge, "")

	fmt.Println("This is a", daystotal, "challenge and you have completed", dayscomplete, "days")

	fmt.Println("Great work") */

	/*s := "你好"
	fmt.Println(len(s)) // 6（UTF-8 编码，每个中文占 3 字节）
	fmt.Println(s[0])   // 228（第一个字节，不是"你"） */

	var user1 User
	user1 = User{
		name:   "王亚鹏",
		age:    21,
		height: 1.76,
	}

	fmt.Println(user1)
	fmt.Println(user1.name)
}
