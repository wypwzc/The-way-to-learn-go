package main

//数组赋值影响复制数组
//修改备份不影响原数组
/*func main() {
	a := [3]int{80, 90, 70}
	b := a
	b[0] = 99
	fmt.Println(a) // Output: [99 90 70]
	fmt.Println(b) // Output: [99 90 70]
}*/

//切片赋值影响原切片
//修改备份影响原切片
/*func main() {
	scores := make([]int, 1, 4)
	scores[0] = 1
	fmt.Println(scores)
	scores = append(scores, 2, 3, 4)
	fmt.Println(scores)
	scores = append(scores, 5)
	fmt.Println(scores)
}*/
