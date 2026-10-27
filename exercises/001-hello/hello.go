//
// 休斯顿，我们遇到问题了。
//
// 你在 WizKidz 公司第一天上班时，了解到软件部门正因公司内部
// 开发的 AI 执行的一次自动化重构出错而陷入困境。
//
// 你的任务是修复这些有问题的程序。
//
// 问题：
// 这个程序原本应该打印 "hello world"，但现在它根本无法编译！
//
// 在 Go 中，函数接受带类型的参数。
//
// 打印字符时常用的类型是 `string` 类型。
// Go 中的字符串通过将文本包裹在双引号 `"` 中来表示。例如："hey there!"
//
// 尝试修复 fmt.Println 函数的参数，然后运行 `go run hello.go`
// 看是否能正常工作。
//
// 如果你正在使用 gopherlings 辅助程序，
// 可以取消注释 "I AM STILL GOING" 来继续
// 解决下一个练习题。

// I AM STILL GOING

package main

import "fmt"

func main() {
	var name string
	count, err := fmt.Scan(&name)
	if err != nil {
		fmt.Println("输入失败")
		return
	}
	fmt.Println(count, name)
}
