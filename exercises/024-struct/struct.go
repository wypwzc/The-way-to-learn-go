// 问题：
// TODO

// I AM STILL GOING

package main

import "fmt"

func main() {
	// 嗯？有人想给这部电影打 10/10 分，
	// 但他们的努力没有成功。
	// 通过修改 setMovie 函数来帮助他们。
	var myMovie movie
	setMovie(myMovie)
	fmt.Println(myMovie)
}

type movie struct {
	title string
	score int
}

// setMovie 实际上是从头创建一个电影。
// 我们应该改变它，使其返回电影类型……并且
// 可能将其名称改为更能描述其功能的名字。
func setMovie(m movie) {
	// 点号索引是我们在 Go 中访问字段的方式。
	m.title = "Everything Everywhere All At Once"
	m.score = 10
	// 电影参数是一个副本，所以函数内的所有更改
	// 在函数退出时都会丢失。我们应该在这里返回一个值。
}
