//
// 问题：
// 在 Go 中，常量正如其名——常量，不可改变的，
// 不可变的。
// Go 支持字符串常量、数值常量和布尔常量。
//
// 在 Go 中不能修改或重新赋值给常量。

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 在 Go 中，变量名必须以字母开头。
	const 20number = 20
	fmt.Println(math.Sin(20number))

	const abc = "abc"
	fmt.Println(abc)

	const sum = abc + "tyk"
	fmt.Println(sum)

	// 这里有问题——尝试给常量赋值会导致编译错误。
	// 与其给已有常量赋值，不如声明一个新的常量！
	sum = sum + "cheeky addition"
	fmt.Println(sum)
}
