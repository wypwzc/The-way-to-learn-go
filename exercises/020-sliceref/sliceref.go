//
// 问题：
// 到目前为止，我们看到函数接收的是参数的副本，
// 无法直接修改它们。但 Go 中的某些类型并非如此。
//
// 我们刚才看到的切片类型实际上持有对底层内存的引用。
// 这意味着切片的副本不会复制切片索引中的内存。
// 如果 Gopher 们不小心，这可能导致意外行为。
//
// 这种引用行为在学习了指针后可能会更清楚。

// I AM STILL GOING

package main

import "fmt"

func main() {
	fibonacci := []int{0, 1, 1, 2, 3, 5, 8, 13}
	fiboAdded := fibonacci
	for i := range fiboAdded {
		fiboAdded[i] += i
	}
	// 哎呀。我们不小心同时修改了两个切片。
	// 尝试使用下面的 clone 函数创建
	// fibonacci 切片的深拷贝，这样对
	// 一个切片的修改就不会影响另一个。
	fmt.Println(fibonacci, fiboAdded)
}

func clone(a []int) []int {
	return append([]int{}, a...)
}
