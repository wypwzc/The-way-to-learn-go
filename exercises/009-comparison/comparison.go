//
// 问题：
// Go 有比较运算符，用于比较两个值并产生一个布尔值。
//
// Go 支持以下比较运算符：
//     a == b   表示 "a 等于 b"
//     a < b    表示 "a 小于 b"
//     a > b    表示 "a 大于 b"
//     a != b   表示 "a 不等于 b"
//     a <= b   表示 "a 小于或等于 b"
//     a >= b   表示 "a 大于或等于 b"

// I AM STILL GOING

package main

import "fmt"

func main() {
	number := 4
	// % 是取模运算符。它返回除法后的余数。
	// 当一个数除以 2 的余数为零时，它是偶数。
	remainder := number % 2


	var is4even bool
	// 这个比较有些问题。
	is4even = remainder = 0

	// 这应该打印 "4 is even:"，后跟 true 或 false，
	// 取决于 4（变量 number）是否真的是偶数。
	fmt.Println(number, "is even:", is4even)
}
