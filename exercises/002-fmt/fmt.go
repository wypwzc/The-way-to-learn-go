//
// 问题：
// 这个程序原本打印 `hello world`，但现在它报编译错误，
// 声称 `fmt` 未定义。这是因为缺少 import 语句。
//
// import 语句以 `import` 关键字开头，后跟
// 用双引号包裹的 import 名称。
//
//  import "fmt"
//
// 当有多个 import 语句时，import 名称在括号内列出。
//
//  import (
//	   "fmt"
//     "strings"
//  )
//
// 很少需要手动修改 import 语句，因为大多数 IDE 可以自动
// 推断你引用的是哪个 import 并在保存时自动添加。

// I AM STILL GOING

package main

// import 语句总是跟在 `package` 声明之后。

func main() {
	fmt.Println("hello world")
}
