package main

import "fmt"

//1方法的定义
/*type Student struct{ Score int }

func (s Student) Passed() bool {
	return s.Score >= 60
}
func main() {
	student := Student{Score: 80}
	fmt.Println(student.Passed())
}*/

//2 修改原记录的方法

type Student struct{ Score int }

func (s *Student) AddBonus(bonus int) {
	s.Score += bonus
}

func main() {
	student := Student{Score: 80}
	student.AddBonus(5)
	fmt.Println(student.Score)
}
