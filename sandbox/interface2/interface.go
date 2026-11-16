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
