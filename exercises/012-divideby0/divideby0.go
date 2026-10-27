//
// 问题：
// 用零除整数是未定义的操作。为了处理这个问题，
// Go 会 `panic` 以避免未定义行为。如果不被 recover() 捕获，
// panic 会导致程序崩溃，关于这点后续会详细说明。
//
// 此外，在 Go 中，声明时未初始化的变量会得到"零值"。
// 例如，声明整数时不赋值，它的值将是零。

// I AM STILL GOING

package main

import "fmt"

func main() {
	var z int
	// 添加一个 if 语句，如果 z 等于零，
	// 则将 z 设置为 7。
	fmt.Println(49/z + 7)
}
