// 问题：
// TODO

// I AM STILL GOING

package main

import (
	"fmt"
	"strings"
)

func main() {
	// 你永远要对用户输入保持警惕。
	// 我们人类就是这样——不可预测且
	// 不喜欢编译器错误。
	// 在处理或显示用户输入之前，通常要"净化"它。
	// 这里我们犯了一个错误：尝试在普通字符串上调用 PrintSafe，
	// 但 PrintSafe 接受一个 `safe` 类型参数，该类型其实只是
	// 一个拥有自己独立类型的字符串。
	// 这阻止了我们打印未净化的输入！要进行净化，
	// 在 unsafeInput 上调用 `sanitize` 来获取一个"safe"字符串！
	unsafeInput := "stupid compiler"
	PrintSafe(unsafeInput)
}

type safe string

func PrintSafe(s safe) {
	fmt.Println(s)
}

func sanitize(possibleBadString string) safe {
	safeString := strings.ReplaceAll(possibleBadString, "stupid", "****")
	return safe(safeString)
}
