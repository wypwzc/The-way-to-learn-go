//
// 问题：
// 切片（slice）是在单个结构中存储相同类型的有序连续元素
// 的最常用方式。在其他语言中，切片被称为数组或列表。
//
// 要初始化切片，可以使用 make 内建函数：
//
//  desiredLength := 23
//  s := make([]int, desiredLength)
//
// Go 中的切片从零开始索引：
//
//  A[0] // 访问 A 的第一个元素
//  A[9] // 访问 A 的第十个元素
//
// Go 确实有数组类型，但这些类型使用非常有限，因为数组的长度
// 也是其类型的一部分。
//

// I AM STILL GOING

package main

func main() {
	// 切片应该有多长才能容纳所有的值？
	beatles := make([]string, ??)

	beatles[0] = "john"
	beatles[1] = "george"
	beatles[2] = "ringo"
	beatles[3] = "paul"

	// 修改这一行，使其打印鼓手的名字。
	fmt.Println(beatles[??])
}
