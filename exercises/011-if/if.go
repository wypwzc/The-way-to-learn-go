// 问题：
// `if` 语句指定一个在条件为 true 时要执行的代码块。
// 它有多种形式。
//
// 简单的 if 语句：
//
//  if x > max {
//  	max = x
//  }
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	const tuesdayDiscount = 0.50 // 50% 折扣！哇！
	isTuesday := 1

	price := 32.00
	// 修复这个条件，使其接受数字：
	if isTuesday {
		price = price * (1 - tuesdayDiscount)
	}
	fmt.Println("The price is", price)
}
