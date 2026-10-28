
## 程序结构

每个Go文件第一行必须声明包名，可执行程序的入口固定是`main`包里的`main`函数：

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

跟Python不同的是，Go里**导入了但没用到的包、声明了但没用到的局部变量都会导致编译失败**，这不是bug，是Go故意逼着你保持代码整洁。`go run main.go`直接运行，`go build`编译成一个不依赖任何运行时的二进制文件。

## 变量、常量与类型

`var`用于声明变量，可以指定类型也可以让编译器推断；`:=`是"短声明"，只能在函数内部使用，且必须同时赋初值：

```go
var name string = "Alice"
var age = 18      // 类型推断为int
score := 90       // 短声明，只能在函数体内用

const Pi = 3.14159

var (
    width  = 100
    height = 200
)
```

Go是静态类型语言，**变量类型一旦确定就不能变**（Python里`x = 1; x = "str"`是合法的，Go里不行）。基本类型包括整数族（`int`、`int8/16/32/64`、`uint`系列）、浮点`float32/float64`、字符串`string`（不可变，UTF-8编码）、布尔`bool`，以及`byte`（即`uint8`，代表一个字节）和`rune`（即`int32`，代表一个Unicode字符）。

类型转换必须显式进行，不存在Python里"int和float混合运算自动提升"的隐式转换：

```go
var i int = 10
var f float64 = float64(i) // 必须显式转换，写 f := i 会编译报错
```

还有一个Python程序员容易忽略的概念：**零值**。Go里声明变量不赋值，会自动得到该类型的零值（`int`是`0`，`string`是`""`，`bool`是`false`，指针/切片/map/channel是`nil`），不会像Python那样访问未赋值变量直接报`NameError`。这意味着很多该判空的地方，新手容易忘记判断。

## 控制流程

Go只有一个循环关键字`for`，靠不同写法模拟`while`和无限循环：

```go
for i := 0; i < 5; i++ {   // 标准三段式
    fmt.Println(i)
}

n := 0
for n < 5 {                 // 当 while 用
    n++
}

for {                        // 无限循环，需要手动 break
    break
}

for index, value := range mySlice {  // 遍历切片/map/字符串
    fmt.Println(index, value)
}
```

`if`可以带一个初始化语句，这是Go里极其常见的写法，尤其常和map查找、错误检查搭配：

```go
if value, ok := myMap["key"]; ok {
    fmt.Println(value)
}
```

`switch`默认不会贯穿到下一个case（不需要写`break`），也可以不带表达式当`if-else`链用：

```go
switch day {
case 1, 7:
    fmt.Println("周末")
default:
    fmt.Println("工作日")
}
```

另外提一句：Go没有三元表达式（`?:`），也没有`**`幂运算符（要用`math.Pow`），这两个是新手最初常找不到对应写法的地方。

## 函数

Go函数最大的特色是**多返回值**和**可变参数**，没有默认参数和关键字参数：

```go
func add(a, b int) int {
    return a + b
}

func divmod(a, b int) (quotient, remainder int) { // 命名返回值
    quotient = a / b
    remainder = a % b
    return // 裸返回，自动带出 quotient 和 remainder
}

func sum(nums ...int) int { // 可变参数
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

多返回值在Python里通常用tuple模拟（`return a, b`），效果类似，但Go里是语言级支持，最常见的用法就是同时返回`(结果, error)`，这个模式会贯穿你写的几乎每一个Go函数。函数也是一等公民，支持闭包：

```go
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}
// c := counter(); c() 返回1，再调用一次返回2……
```

## 复合数据类型：数组、切片、map、struct

~~~go
//数组：
//数组赋值影响复制数组
//修改备份不影响原数组
//不能把这条结论随便套给切片
a := [3]int{80,90,70}
b := a
b[0] = 99

~~~

~~~go
//切片:[]里不写数字就是切片
//切片可以理解为看向底层元素的一个窗口
//切片范围左边包含右边不包含
func main() {
	scores := []int{80, 90, 70}
    part := scores[1:3]
	part[0] = 95
	fmt.Println(scores)
	fmt.Println(part)
    
运行结果：
    [80 95 70]
    [95 70]

1.切片是对底层数组的引用
	Go 语言中的切片本身并不存储数据，它是一个轻量级的数据结构，内部包含三个字段：一个指向底层数组的指针、切片的长度（Len）和切片的容量（Cap）。
	当执行 scores := []int{80, 90, 70} 时，Go 在底层创建了一个包含 [80, 90, 70] 的数组，scores 切片指向这个数组。
2.共享底层数组
    执行 part := scores[1:3] 时，part 切片截取了 scores 从索引 1（包含）到 3（不包含）的元素。
    此时，part 包含的元素是 [90, 70]，但更重要的是，part 和 scores 共享同一个底层数组。part 的起始指针指向底层数组中原来 90 所在的位置。
3.修改互相影响
    执行 part[0] = 95 时，实际上修改的是底层数组中索引为 1 的那个元素（原本是 90，现在被改成了 95）。
    因为 scores 和 part 都引用了同一个底层数组，所以当通过 scores 打印时，索引 1 的位置已经变成了 95，最终输出 [80, 95, 70]。
~~~



数组长度是类型的一部分（`[3]int`和`[5]int`是不同类型），实际开发几乎不直接用数组，99%场景用**切片（slice）**。这是整个基础阶段最容易踩坑的地方：切片本质是对底层数组的一个"窗口"（指针+长度+容量），切片操作不是拷贝，而是共享底层数据：

```go
s1 := []int{1, 2, 3, 4, 5}
s2 := s1[1:3]      // s2 是 [2, 3]，但和s1共享底层数组
s2[0] = 100        // 直接修改了s1！s1 变成 [1, 100, 3, 4, 5]

s3 := append(s2, 999) // 如果s2容量够用，会继续覆盖s1后面的元素
// 此时 s1 变成 [1, 100, 3, 999, 5]，s3 是 [100, 3, 999]
```

这和Python里`list[1:3]`会生成一个全新list完全不同——Python的切片是拷贝，Go的切片是视图。不确定是否安全的时候，用`copy()`显式复制一份。

`map`用法和Python的`dict`很像，但访问不存在的key不会报错（而是返回零值），所以一定要用`v, ok := m[key]`的写法判断是否存在；另外Go的map遍历顺序是随机的，不像Python 3.7+的dict会保留插入顺序：

```go
m := make(map[string]int)
m["a"] = 1
v, ok := m["b"]   // ok为false，对应Python的 m.get("b") 或 "b" in m
delete(m, "a")
```

`struct`是Go里组织数据的方式，没有继承，用"嵌入"（embedding）实现组合：

```go
type Animal struct{ Name string }
type Dog struct {
    Animal      // 匿名嵌入，自动获得 Animal 的字段
    Breed string
}
d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Husky"}
fmt.Println(d.Name) // 可以直接访问，因为嵌入了Animal
```

## 指针与方法

Go函数默认**传值**（拷贝一份），如果要在函数内修改外部变量，需要传指针：

```go
func modify(n *int) {
    *n = 100
}

func main() {
    x := 10
    modify(&x)
    fmt.Println(x) // 100
}
```

Python里所有对象传参本质上都是传引用，所以习惯了"在函数里改了对象的属性外面就能看到"，到Go里第一次发现"明明调用了方法，值却没变"，往往就是值receiver和指针receiver搞混了：

```go
type Rectangle struct{ Width, Height float64 }

func (r Rectangle) Area() float64 {       // 值receiver，操作的是副本
    return r.Width * r.Height
}

func (r *Rectangle) Scale(factor float64) { // 指针receiver，能修改原对象
    r.Width *= factor
    r.Height *= factor
}

rect := Rectangle{Width: 10, Height: 5}
rect.Scale(2)            // Go自动处理成 (&rect).Scale(2)
fmt.Println(rect.Area()) // 200，因为Scale真的改了rect本身
```

原则很简单：方法需要修改receiver、或者struct较大想避免拷贝开销时用指针receiver，否则用值receiver；同一类型的方法尽量保持receiver类型一致。

## 包与可见性

可见性规则是Go特有的：**标识符首字母大写就是导出的（包外可访问），小写就是包内私有**，没有`public`/`private`关键字，编译器强制检查，不像Python的下划线前缀只是君子协定：

```go
// calculator/calc.go
package calculator

func Add(a, b int) int      { return a + b } // 大写，可被外部包调用
func subtract(a, b int) int { return a - b } // 小写，仅包内可用
// main.go
package main

import (
    "fmt"
    "myproject/calculator" // 路径基于 go.mod 里声明的 module 名
)

func main() {
    fmt.Println(calculator.Add(1, 2)) // OK
    // calculator.subtract(1, 2)      // 编译错误：未导出
}
```

## 错误处理的基本样子（先预告一下）

虽然系统讲错误处理是阶段二的内容，但从写第一行Go代码起你就会天天碰到它，这里先打个底：Go没有`try/except`，错误就是一个普通的返回值，调用者必须显式检查：

```go
result, err := strconv.Atoi("123")
if err != nil {
    fmt.Println("转换失败:", err)
    return
}
fmt.Println(result)
```

"调用函数 → 检查err → 处理或继续"这个模式会出现在你写的几乎每个Go函数里，先混个脸熟，阶段二再展开自定义错误、错误包装这些细节。

以上就是阶段一对应的全部基础语法。建议不要一次性看完就去写项目，而是照着上面每个代码块自己敲一遍、改一改参数看看输出，尤其是切片共享底层数组那个例子，亲手跑一次比看十遍文字理解更深刻。