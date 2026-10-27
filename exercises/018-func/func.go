//
// 问题：
// 函数是 Go 程序执行的核心。
// 你可能已经注意到，你一直在名为 `main` 的函数体内编写
// 你的 Go 程序。
//
//  func main() {
//  	// 你的代码写在这里！
//  }
//
// `main` 包中的 `main` 函数定义了一个可执行的 Go 程序。
// 它在程序运行时被自动调用。
//
// 你也可以在 `main` 之外定义可执行的代码片段！
// 下面是一个将两个整数相加并返回结果的函数。
//
//  func add(a int, b int) int {
//  	c := a + b
//  	return c
//  }
//
// 然后可以在代码中轻松使用这个函数！
//
//  sum := add(1, 2)
//
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	sq := square(3)
	fmt.Println(sq)
}

// 哎呀！有人想计算一个数的平方
// 并返回结果，但看起来
// 有些不对。
// 记住，函数中括号后面的类型
// 就是返回类型。当一个函数定义了
// 返回值时，还必须记得使用 `return` 关键字
// 指定要返回哪个值。
func square(a int) int {
	a * a
}
