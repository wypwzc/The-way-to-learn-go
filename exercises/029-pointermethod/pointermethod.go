// 问题：
// 就像函数可以有指针参数一样，方法也可以有指针接收者。
//
// 使用指针接收者有两个原因。
//
// 第一个是方法可以修改其接收者指向的值。
//
// 第二个是避免在每次方法调用时复制值。
// 例如，如果接收者是一个大型结构体，这会更高效。
//
// 一般来说，给定类型上的所有方法应该要么都是值接收者，要么都是指针接收者，
// 但不要混合使用两者。我们很快就会看到原因。

// I AM STILL GOING

package main

import "fmt"

func main() {
	var id universalID
	id.Set(20)
	// 看起来我们打印的是 universalID 的零值 (0)。
	// 我们需要 id.Set 能够实现设置 id 为 Set 参数的效果，
	// 可以通过更改 Set 方法中的代码来实现。
	fmt.Println(id.Get())
}

type universalID int

// Set 将接收者 universalID 设置为 newID。
func (uid universalID) Set(newID int) {
	uid = universalID(newID)
}

func (uid universalID) Get() int {
	return int(uid)
}
