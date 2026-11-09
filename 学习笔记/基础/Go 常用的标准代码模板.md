### 1. 数组 (Array)

数组是**固定长度**的同类型数据集合。

```Go
// 1. 声明后赋默认值（元素全为 0）
var arr1 [3]int 

// 2. 声明并初始化
arr2 := [3]int{1, 2, 3}

// 3. 让编译器自动计算长度 (...)
arr3 := [...]string{"Go", "Python", "Java"}
```

### 2. 切片 (Slice)

切片是动态的、灵活的“动态数组”，开发中最常用。

[]里不写数字就是切片

```Go
// 1. 声明空切片（nil）
var s1 []int

// 2. 字面量初始化
s2 := []int{1, 2, 3, 4, 5}

// 3. 使用 make() 创建（长度为 3，容量为 5）
s3 := make([]int, 3, 5)

// 4. 追加元素
s3 = append(s3, 10, 20)
```

### 3. 映射 (Map)

Map 是键值对（Key-Value）结构。

```Go
// 1. 声明（此时为 nil，不能直接赋值，需先 make）
var m1 map[string]int

// 2. 使用 make() 初始化
m2 := make(map[string]int)

// 3. 声明并直接初始化字面量
m3 := map[string]int{
    "apple":  5,
    "banana": 3,
}

// 4. 增删改查
m3["orange"] = 10 // 增/改
delete(m3, "apple") // 删
val, exists := m3["banana"] // 查（判断是否存在）
```

### 4. 函数 (Function)

Go 函数支持多返回值。

```Go
func 函数名[T 约束类型](参数名 参数类型) 返回值类型 {    // 函数体逻辑}
// 基本函数模板（带参数和多个返回值）
func greet(name string) string {
	return "Hello," + name
}

func main() {
	a := greet("xiaowang")
	println(a)
}
```

### 5. 流程控制结构 (Control Structures)

这里指代常见的 `if`、`for` 和 `switch` 结构：

```Go
// 条件判断 if-else
if x := 10; x > 5 { // 可以在 if 中声明临时变量
    // ...
} else {
    // ...
}

// 循环 for（Go 中没有 while，统一用 for）
for i := 0; i < 10; i++ {
    // ...
}

// 遍历切片/数组/Map
for index, value := range s2 {
    // ...
}
```

### 6. 结构体 (Struct)

用于自定义聚合数据类型。

```Go
// 1. 定义结构体
type Person struct {
    Name string
    Age  int
    Email string
}

// 2. 初始化结构体
// 方式 A：按字段名初始化（推荐）
p1 := Person{
    Name:  "张三",
    Age:   18,
    Email: "zhangsan@example.com",
}

// 方式 B：按顺序初始化
p2 := Person{"李四", 20, "lisi@example.com"}
```

### 7. 方法 (Method)

方法是作用在特定结构体（接收者 Receiver）上的函数。

```go
// 定义一个结构体
type Rectangle struct {
    Width, Height float64
}

// 定义方法（使用指针接收者，可以修改结构体内部的值，或避免大对象拷贝）
func (r *Rectangle) Area() float64 {
    return r.Width * r.Height
}

// 调用方式
// rect := Rectangle{Width: 10, Height: 5}
// area := rect.Area() // 输出 50
```

### 8. 接口

```Go
// 定义一个接口
type 接口名 interface {
    方法名1(参数列表) 返回值类型
    方法名2(参数列表) (返回值1, 返回值2)
}
```

完整的使用模板（定义、实现、调用）

Go 语言的接口实现是**隐式的**——你不需要写 `implements` 关键字，只要某个类型实现了接口中定义的所有方法，它就自动实现了该接口。

```Go
package main

import "fmt"

// 1. 定义接口
type Greeter interface {
    Greet() string
}

// 2. 定义一个结构体
type User struct {
    Name string
}

// 3. 为结构体实现接口中的方法（隐式实现）
func (u User) Greet() string {
    return "你好，" + u.Name
}

// 4. 编写一个接收接口类型作为参数的函数（利用多态）
func DoGreet(g Greeter) {
    fmt.Println(g.Greet())
}

func main() {
    // 5. 初始化对象并传入接口
    u := User{Name: "张三"}
    DoGreet(u) // 输出: 你好，张三
}
```