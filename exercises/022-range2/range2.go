// 问题：
// TODO

// I AM STILL GOING

package main

import "fmt"

func main() {
	fruitScores := map[string]int{
		"apple":     5,
		"banana":    8,
		"tangerine": 10,
	}
	// 为什么这不能编译！？
	// Go 对未使用的变量很挑剔。
	// 在这个特定案例中，我们可以大幅减少代码量。
	// 在这个 for 块内使用 v 来替换 map 访问，
	// 并打印出评分高于 7 的水果。
	for k, v := range fruitScores {
		if fruitScores[k] > 7 {
			fmt.Println(k, " has score ", fruitScores[k])
		}
	}
}
