// 问题：
// TODO
// https://go.dev/tour/basics/13
// 表达式 T(v) 将值 v 转换为类型 T。

// I AM STILL GOING

package main

import "fmt"

func main() {
	var three int = 3
	var oneThird float64 = 1.0 / 3.0
	var twoPlusThree int = 2 + three
	// 这两个变量是不同类型的。要在变量之间执行操作，
	// 它们必须是相同类型！将 Println 中一个变量转换为
	// 与另一个匹配的类型，以便打印出 1/3+2+3 的结果。
	// 不要添加新行或更改上面的行！
	fmt.Println(oneThird + twoPlusThree)
}
