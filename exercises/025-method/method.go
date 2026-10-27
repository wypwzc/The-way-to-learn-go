//
// 问题：
// Go 没有类（class）。但是，你可以在类型上定义方法。
//
// 方法是一个带有特殊接收者参数的函数。
// 接收者出现在它自己的参数列表中，位于 func 关键字和方法名之间。
// 在这个例子中，Info 方法有一个名为 b 的 book 类型接收者。
//
// 请记住，方法就是函数。下面的方法可以重写为
// 一个常规函数，功能不变：
//
//  func Info(b book) string {
//		// 相同的代码
//  }

// I AM STILL GOING

package main

import (
	"fmt"
	"strconv"
)

func main() {
	checkedOut := book{
		title:  "Los Hijos de los Días",
		author: "Eduardo Galeano",
		year:   2012,
	}
	// 嗯，这个程序至少做对了一件事……
	// 它打印出了结构体的未格式化内容。
	// 但我们想要打印书籍的格式化信息——
	// 可以通过调用某个返回书籍格式化信息的方法来实现。
	fmt.Println(checkedOut)
}

type book struct {
	title  string
	author string
	year   int
}

// 值得注意的是，我们也可以通过简单地将此方法重命名为"String"来解决这个练习。
// fmt 包会检查所有参数是否实现了 `String` 方法，
// 并使用它来打印参数，覆盖默认的打印算法。
// 参见 fmt.Stringer 类型。
func (b book) Info() string {
	yearString := strconv.Itoa(b.year) // 将年份转换为字符串！
	return "\"" + b.title + "\" by " + b.author + " (" + yearString + ")"
}
