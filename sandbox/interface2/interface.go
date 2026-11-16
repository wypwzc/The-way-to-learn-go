/*接口例子
package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}

type Student struct {
	Human  // 匿名字段Human
	school string
	loan   float32 // 加上这行
}

// Human对象实现Sayhi方法
func (h *Human) SayHi() {
	fmt.Printf("Hi, I am %s, you can call me on %s\n", h.name, h.phone)
}

// Human对象实现Sing方法
func (h *Human) Sing(lyrics string) {
	fmt.Println("La la, la la la, la la la la la...", lyrics)
}

func main() {
	s := Student{
		Human: Human{
			name:  "张三",
			age:   20,
			phone: "138-1234-5678",
		},
		school: "某某大学",
		loan:   1000.0, // 初始化 loan，顺便消除警告
	}

	fmt.Println("--- 调用 Sing 方法 ---")
	s.Sing("《晴天》")

	fmt.Println("\n--- 调用 SayHi 方法 ---")
	s.SayHi()

	fmt.Println("\n--- 访问底层字段 ---")
	fmt.Printf("姓名: %s\n", s.name)
	fmt.Printf("年龄: %d\n", s.age)
	fmt.Printf("学校: %s\n", s.school)
	fmt.Printf("贷款: %.2f\n", s.loan) // 读取它，彻底消除 unused write 警告
}
*/

package main

import "fmt"

type Human struct {
	name  string
	age   int
	phone string
}
type Student struct {
	Human  //匿名字段
	school string
	loan   float32
}
type Employee struct {
	Human   //匿名字段
	company string
	money   float32
}

//Human实现SayHi方法
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

//Human实现Sing方法
func (h Human) Sing(lyrics string) {
	fmt.Println("La la la la...", lyrics)
}

//Employee重载Human的SayHi方法
func (e Employee) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s\n", e.name,
		e.company, e.phone)
}

// Interface Men被Human,Student和Employee实现
// 因为这三个类型都实现了这两个方法
type Men interface {
	SayHi()
	Sing(lyrics string)
}

func main() {
	mike := Student{Human{"Mike", 25, "222-222-XXX"}, "MIT", 0.00}
	paul := Student{Human{"Paul", 26, "111-222-XXX"}, "Harvard", 100}
	sam := Employee{Human{"Sam", 36, "444-222-XXX"}, "Golang Inc.", 1000}
	tom := Employee{Human{"Tom", 37, "222-444-XXX"}, "Things Ltd.", 5000}
	//定义Men类型的变量i
	var i Men
	//i能存储Student
	i = mike
	fmt.Println("This is Mike, a Student:")
	i.SayHi()
	i.Sing("November rain")
	//i也能存储Employee
	i = tom
	fmt.Println("This is tom, an Employee:")
	i.SayHi()
	i.Sing("Born to be wild")
	//定义了slice Men
	fmt.Println("Let's use a slice of Men and see what happens")
	x := make([]Men, 3)
	//这三个都是不同类型的元素，但是他们实现了interface同一个接口
	x[0], x[1], x[2] = paul, sam, mike
	for _, value := range x {
		value.SayHi()
	}
}
