//
// 问题：
// 下面是一个曾经能打印值的程序。看起来
// 重构的人不知道 fmt.Println 接受多个参数，
// 通过逗号分隔不同类型的参数。
//
//  fmt.Println("hello", "world", 99)
//
// 上面的程序打印 "hello world 99"。可变参数函数接受
// 任意数量的参数。

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 了解如何在函数调用中用逗号分隔参数后，修复
	// 下面这行代码使其打印 "go 2 true"
	fmt.Println("go" 2 true)
}
