// 问题：
// Go 也支持 if-else 语句
//
// 带有 else 分支的 if 语句，当条件为 false 时执行 else 分支：
//
//  if x > max {
//  	max = x
//  } else {
//  	x = max
//  }
//
// 带有 else if 和/或前置语句的 if 语句
//
//  if x := f(); x > max {
//  	max = x
//  } else if x == max {
//  	max = max + 1
//  } else {
//  	x = max
//  }
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	number := 7
	if number > 0 {
		fmt.Println("number is greater than 0")
	} {
		fmt.Println("number is less than 0")
	}

	// 请确保条件语句是正确的。
	if number = 0 {
		fmt.Println("number is zero")
	} else if number % 2 != 0 {
		fmt.Println("number is even")
	} else {
		fmt.Println("number is odd")
	}
}
