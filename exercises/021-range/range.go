// 问题：
// "for" 语句可以使用 "range" 子句来迭代
// 切片、map 和其他类型中的所有值。
//
//  for i := range slice {
//  	slice[i] = f()
//  }
//
// 上面的代码将切片中的所有项设置为
// 调用函数 f 返回的值。
//
// 也可以为 range 添加第二个值。
//
//  for i, v := range slice {
//  	slice[i] = v * v
//  }
//
// 上面的代码对切片中的所有值求平方。
// 请记住，range 子句返回的第一个值
// 是切片的索引。range 子句的第二个
// 值就是切片中的值。
//

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 该死，这个程序本该将切片中的值翻倍，
	// 但它做的事情完全不是这样。
	//
	// 注意 range 子句返回的索引和值的顺序。
	slice := []int{1, 2, 3, 4}
	for i, v := range slice {
		slice[v] = 2 * i
	}
	fmt.Println(slice)
}
