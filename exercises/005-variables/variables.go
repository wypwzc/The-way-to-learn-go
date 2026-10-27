//
// https://gobyexample.com/variables
//
// 问题：
// 在 Go 中，变量声明时带有类型，以便编译器
// 检查函数调用的正确性。
//
// 关键字 `var` 用于声明一个或多个变量。
// 以下代码声明了一个名为 `a`、类型为 `int` 的变量。
//
//  var a int
//
// 然后你可以在声明语句下方多次给新建的值赋值：
//
//  a = 2
//
// 你也可以在同一行声明并赋值，效果相同：
//
// var a int = 2
//
// 在一行中声明多个变量：
//
//  var a, b string = "hello", "world"
//
// 在 Go 中，常见的是省略变量声明中的类型：
//
//  var a = "look ma, no type!"
//
// 最短（也可能是最常见）的声明和初始化方式
// 是使用简写 `:=`：
//
//  a := "as short as you can"
//  b := 4
//
// 在 Go 中，声明时未初始化的变量会得到"零值"。
// 例如，字符串会是空字符串 ""，数字会是 0。
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 赋值给 `a` 的值有一个与 `a` 的类型
	// 不匹配的类型。修复这一行，
	// 使得下面的打印输出 "123"。
	var a string

	a = 123
	fmt.Println(a)

	// 我们试图声明并初始化变量 `b` 和 `c`，
	// 但有一个语法错误。只需要加 2 个字符就能解决！
	var b c int = 2*1 2*2
	fmt.Println(b, c)

	// 有人尝试使用简写初始化但搞错了。
	d = "pear"
	fmt.Println(d)
}
