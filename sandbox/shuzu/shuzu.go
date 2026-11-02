package main

import "fmt"

//数组赋值影响复制数组
//修改备份不影响原数组
func main() {
	a := [3]int{80, 90, 70}
	b := a
	b[0] = 99
	fmt.Println(a) // Output: [99 90 70]
	fmt.Println(b) // Output: [99 90 70]
}
