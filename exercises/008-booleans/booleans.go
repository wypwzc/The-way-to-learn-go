//
// 问题：
// 布尔类型定义了以下逻辑运算符：
//
//  &&    条件与（AND）    p && q  仅当 p 和 q 都为 true 时结果为 true。
//  ||    条件或（OR）     p || q  如果 p 或 q 任一为 true 则结果为 true。
//  !     非（NOT）        !p      如果 p 为 false 则结果为 true。
//
// 你也可以在布尔类型上使用相等检查 `==` 和不相等检查 `!=`。
// 后者模拟了异或操作（XOR）。
//
//  !=   比较 不等于（NOT EQUAL）  p != q  如果 p 不等于 q 则结果为 true
//  ==   比较 等于（EQUAL）        p == q  如果 p 等于 q 则结果为 true
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	a := true
	b := false
	// 修复下面这行，使其打印出 a OR b 的结果。
	fmt.Println(a + b)

	// 修复下面这行，使其打印出 a AND b 的结果。
	fmt.Println(a & b)
}
