//
// 问题：
// 函数可以有多个参数和多个返回值。
// 要创建具有多个返回值的函数，请将
// 返回类型放在括号中。
//
//  func fruits() (string, string) {
//  	return "banana", "tangerine"
//  }
//
// 如果两个或更多连续参数具有相同类型，可以省略
// 类型名称，直到最后一个参数。
//
//  func divmod(a, b int) (div, mod int) {
//  	div = a/b
//  	mod = a%b
//  	return div, mod
//  }
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	fmt.Println(return1And2())
}

// 这个函数应该返回两个值：1 和 2。
// 通过添加两个带名字的返回参数（类似于问题描述中的最后一个示例）
// 来修复函数签名（包含 `func` 关键字的那行）。
func return1And2() {
	a = 1
	b = 2
	return a, b
}
