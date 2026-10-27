// 问题：
// Map 类型是 Go 的关联数据类型，在其他语言中通常称为
// hash 或 dict（字典）。
//
// Map 包含一组无序的同一类型的元素（称为元素类型），
// 由另一组唯一键索引。
//
// 在 Go 中，map 在使用前必须初始化。为此我们使用
// make 内建函数：
//
//  myMap := make(map[int]int)
//
// 在 map 中查找不存在的值时，将返回该元素类型的零值。

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 一个键为字符串类型、元素为整数类型的 map。
	var age map[string]int

	// 如果 map 在访问键之前没有正确初始化，
	// 程序将会崩溃。
	// 修复上面的声明！
	age["paul"] = 32
	age["miguela"] = 27
	age["sarah"] = 78

	fmt.Println(age)

	// 访问 map 时要小心键——它必须与你
	// 正在查找的键完全匹配。
	// 如果找不到键，访问将返回
	// 该元素类型的零值。
	fmt.Println("Sarah's Age:", age["Sarah"])
}
