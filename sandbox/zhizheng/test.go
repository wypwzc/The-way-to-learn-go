package main

import "fmt"

/*func main() {
	a := 10
	b := &a
	*b = 20
	fmt.Println(a, b, *b)
}*/
//输出结果20 0x1c0c0c0 20

//go语言中函数参数传递是值传递，函数内部对参数的修改不会影响到外部变量的值。
func addLocal(score int)     { score += 5 }
func addOriginal(score *int) { *score += 5 }
func main() {
	score := 80
	addLocal(score)
	fmt.Println(score)
	addOriginal(&score)
	fmt.Println(score)
}
