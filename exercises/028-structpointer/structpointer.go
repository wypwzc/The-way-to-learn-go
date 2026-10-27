// 问题：
// TODO

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 和之前看到的程序相同，但现在
	// 我们想用指针来解决它。
	// 如果我们有一个指向结构体的指针，
	// 修改结构体的字段将反映在
	// 指针引用的结构体上。
	var myMovie movie
	setMovie(myMovie)
	fmt.Println(myMovie)
}

type movie struct {
	title string
	score int
}

// 你可能只需要修改这里的一行。
func setMovie(m movie) {
	m.title = "Blue Planet"
	m.score = 10
}
