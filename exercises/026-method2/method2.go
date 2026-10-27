// 问题：
// TODO
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 注意，与结构体、map 和切片等复合字面量不同，
	// 我们不能用花括号 {} 来初始化 XORer，
	// 我们实际上是从 bool 原始类型到 XORer 进行类型转换。
	// 这之所以有效，是因为 XORer 的底层类型是 bool。
	True := XORer(true)
	False := XORer(false)
	fmt.Println("true ^ true =", True.xor(true))
	fmt.Println("false ^ true =", False.xor(true))
	fmt.Println("false ^ false =", False.xor(false))
}

// 许多 Go 新手会惊讶地发现，你可以在任何用户定义的类型上定义方法，
// 包括像本例这样由底层 bool 类型组成的类型！
type XORer bool

// Go 中没有 XOR 操作，我们使用不等运算符在 bool 类型的两个值之间执行异或操作。
// 这段代码有一个错误，因为操作是在两个不同的类型之间进行的。
// 这两个值应该是相同类型才能执行不等操作！
func (a XORer) xor(b bool) bool {
	return a != b
}
