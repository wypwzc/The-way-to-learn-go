package main

import "fmt"

//函数函数基本用法返回值
/*func greet(name string) string {
	return "Hello," + name
}

func main() {
	a := greet("xiaowang")
	println(a)
}
*/

//函数函数基本用法返回值多值
/*func divide(total, count int) (int, bool) {
	if count == 0 {
		return 0, false
	}
	return total / count, true
}
func main() {
	result, ok := divide(10, 2)
	fmt.Println(result, ok)
	result1, ok1 := divide(10, 0)
	fmt.Println(result1, ok1)
}
*/

//defer在函数退出时执行，常用于资源释放
func main() {
	defer fmt.Println("先登记")
	defer fmt.Println("后登记")
	fmt.Println("处理中")
}
