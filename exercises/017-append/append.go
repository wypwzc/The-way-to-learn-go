// 问题：
// append 是一个内建函数，用于构造运行时大小不确定的切片。
//
// 要向现有切片追加一个元素，可以使用 append 函数：
//  short := []string{"first"}
//  longer := append(short, "second")
//
// 结果存储在 `longer` 中，是一个包含两个元素的切片，"first" 后跟
// "second"。
//
// 一个常见模式是追加到同一个切片：
//
//  elements = append(elements, "an item")
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	var evens []int
	for i := 1; i <= 20; i++ {
		if i%2 == 0 {
			// 了解了切片的基本机制后，
			// 你会如何构建一个包含 1 到 20 之间所有偶数的切片？
			evens = ??
		}
	}
	fmt.Println(evens)
}
