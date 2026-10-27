// 问题：
// https://gobyexample.com/pointers
// Go 支持指针，允许通过引用将值传递给函数和方法。
//
// 这意味着带有指针参数的函数将能够修改指针指向的值。
// 在 Go 中，要获取值的引用，我们使用 "&" 运算符：
//
//  var a int = 2
//  ref := &a
//
// ref 将是 *int 类型，也就是说 ref 是一个整数的引用。
// 要访问底层值，我们必须使用 * 运算符解引用指针，
// 该运算符放在指针值前面：
//
//  fmt.Println(*ref) // 打印 2
//  fmt.Println(ref)  // 打印指针在内存中的地址
//
// 要给指针赋值，我们也需要解引用：
//
//  *ref = 1337
//  fmt.Println(*ref) // 打印 1337
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 这是目前为止最难的练习，准备好迎接挑战吧！
	// 我们需要修改这个程序中的至少 3 行才能达到
	// 期望的结果：zeroIt 应该修改 value，而不需要
	// 返回值参数。
	value := 2
	zeroIt(value) // 我们可能需要在这里通过引用传递 `value`。
	fmt.Println(value)
}

// 我们想从函数内部修改 p，所以需要指针的帮助！
// 更改参数类型，使 zeroIt 接受一个指针，
// 然后修改函数体，将零赋值给 p。
func zeroIt(p int) {
	p = 0
}
