# 数组与切片

### 数组：

~~~go
//数组：
//数组赋值影响复制数组
//修改备份不影响原数组
//不能把这条结论随便套给切片
a := [3]int{80,90,70}
b := a
b[0] = 99

~~~

### Slice切片:

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
    
    核心规则总结
引用传递特性：对切片元素的修改（如 slice[i] = val），会直接反应到原切片和所有共享该底层数组的其他切片上。

长度与容量：切片操作（如 [1:3]）不会复制数据，只是创建了一个新的切片视图，指向原数组的特定区间。如果想要完全独立的数据副本，应该使用 copy() 函数复制一份。
~~~



#### 创建切片

~~~go
func main() {
	scores := make([]int, 2, 4)
    scores[0], scores[1] = 80, 90
    scores = append(scores, 70)
    fmt.Println(scores)
	fmt.Println(len(scores), cap(scores))
}

1.make([]int, 2, 4) 的初始化
	该语句创建了一个切片，其长度（Len）为 2，容量（Cap）为 4。
	这意味着底层数组分配了 4 个元素的空间（初始值全为 0），但当前切片只能直接访问前 2 个元素。
2.索引赋值与长度限制
	scores[0], scores[1] = 80, 90：直接通过索引给前两个元素赋值。此时切片内容变为 [80, 90]。
	规则：只能通过切片索引访问或修改 < len 范围内的元素。如果试图直接写 scores[2] = 70 会触发运行时恐慌（panic: index out of range），因为索引 2 超出了当前的长度 2。
3.append 函数与容量复用
	scores = append(scores, 70)：向切片追加元素。
	核心规则：当 append 后的元素总数未超过容量（len + 1 <= cap）时，Go 语言不会重新分配底层数组，而是直接将新元素写入底层数组中空闲的位置（即索引 2 的位置，原本为 0），并将切片的长度加 1。
	在本例中：追加前长度为 2，容量为 4；追加一个元素后，长度变为 3，而容量仍然保持为 4。

核心规则总结
长度（Len）：切片当前包含的元素个数，决定了可以直接用索引访问的范围。

容量（Cap）：切片底层数组从切片起始位置到底层数组末尾的总长度，决定了在不重新分配内存（不扩容）的前提下最多可以存多少个元素。

append 扩容临界点：只要 len 小于 cap，append 就直接在原底层数组上操作，不会发生底层数组的内存重新分配。
~~~



### for range 循环

~~~go
//for range 循环：
func main() {
    scores:= []int{80, 90, 70}
    for index,  score := range scores {
        fmt.Println(index, score)
    }
}


//输出结果：
0 80
1 90
2 70
~~~

**`range` 关键字的返回值**

- 在对切片（Slice）或数组使用 `range` 时，它每次迭代会返回两个值：
  - **第一个值**：当前元素的**索引（Index）**，从 `0` 开始递增。
  - **第二个值**：当前索引处元素的**值（Value）\**的\**副本**。

**可同时给出索引和值**

- 语法形式为 `for index, value := range slice`。
- 如果你只需要索引，甚至可以省略后面的逗号和值，写成 `for index := range slice`。

**下划线明确忽略某个结果**

- Go 语言对未使用的局部变量有严格的编译检查（声明了却不用会直接**编译报错**）。
- 如果你只想要值而不需要索引，就可以用下划线忽略索引：`for _, value := range slice`。

**元素值是副本（与原元素地址不同）**

- 迭代时产生的 `value` 变量是一个全局复用的临时变量，它只保存当前元素的值拷贝。
- 正因为它是独立的副本，所以直接修改 `value` 无法影响原切片；如果想要修改原数据，必须通过索引（如 `scores[index] = ...`）去操作底层数组。

**值拷贝机制（Value Copy）**

- 在 `for index, score := range scores` 中，变量 `score` 并不是原切片中元素的引用（指针），而是该元素的一个**值拷贝**。
- **这意味着**：如果在循环体内对 `score` 变量进行重新赋值（例如 `score = 100`），**只会改变局部变量 `score` 的值，原切片 `scores` 里的内容不会受到任何影响**。

#### 如何在循环中修改切片元素？

如果你想在循环中修改切片里的实际数据，千万不要直接修改 `score` 变量，而应该**通过索引去修改原切片**：

```Go
func main() {
    scores := []int{80, 90, 70}
    
    // 正确的修改方式：通过索引修改原切片
    for index := range scores {
        scores[index] += 5 // 给每个分数加 5 分
    }
    
    fmt.Println(scores) // 输出: [85 95 75]
}
```





### 通过索引修改元素

#### make创建独立底层数组

#### copy深拷贝

~~~go
func main() {
    scores := []int{80, 90}
    for i := range scores { 
        scores[i] += 5 
    }
    backup := make([]int, len(scores))
    count := copy(backup, scores)
    backup[0] = 99
    fmt.Println(scores, backup, count)
}


//输出结果：
[85 95] [99 95] 2
~~~

这段代码主要展示了 Go 语言中切片的**通过索引修改元素**、**`make` 创建独立底层数组**以及 **`copy` 深拷贝**的核心机制：

1. **通过索引在循环中修改切片**
   - `for i := range scores { scores[i] += 5 }`：利用 `range` 循环获取切片的索引 `i`，并通过索引直接修改原切片 `scores` 中的元素。修改后，`scores` 的元素变为了 `[85, 95]`。   
2. **`make` 分配独立的底层数组**
   - `backup := make([]int, len(scores))`：使用 `make` 函数创建了一个**创建一个名字叫 `backup`、类型为 `[]int`、且长度与 `scores` 完全相同的新切片**。
   - **机制**：`make` 会为新切片分配一块**全新的、互不干扰的底层数组**，而不是共享引用。   
3. **`copy` 函数进行深拷贝与返回值**
   - `count := copy(backup, scores)`：将 `scores` 中的元素完整复制到 `backup` 中。   
   - **机制**：`copy(dst, src)` 会把源切片的内容复制到目标切片，并**返回实际复制的元素个数**（此处 `count` 的值为 `2`）。此时 `backup` 的内容更新为 `[85, 95]`。   
   - `copy` 的返回值是**实际复制的元素个数**
4. **独立修改互不影响**
   - `backup[0] = 99`：将 `backup` 的第一个元素修改为 `99`。   
   - **机制**：由于 `backup` 和 `scores` 拥有完全独立的底层数组（实现了深拷贝），因此对 `backup` 的修改**不会影响** `scores`。   
   - 最终 `fmt.Println(scores, backup, count)` 的输出结果为：`[85, 95] [99, 95] 2`。   



### **底层数组共享机制**以及**长度（Length）计算规则**：

~~~go
func main() {
    scores := []int{80, 90, 70}
    part := scores[1:3]
    part[0] = 95
    fmt.Println(scores[1])
    fmt.Println(len(part))
}


//输出结果
95
2
~~~

**截取与共享底层数组**：`part := scores[1:3]` 从索引 `1`（包含）截取到索引 `3`（不包含）。这使得 `part` 包含了原切片中的 `90` 和 `70`，并且`part` 与 `scores` 共享同一个底层数组。   

**修改共享数据**：`part[0] = 95` 修改了 `part` 的第一个元素。因为 `part[0]` 指向底层数组中原来 `90` 所在的位置，所以该修改会直接反映到共享的底层数组中。   

**打印修改后的元素**：`fmt.Println(scores[1])` 访问 `scores` 索引为 `1` 的元素，由于已经被修改，因此会输出 `95`。   

**获取切片长度**：`fmt.Println(len(part))` 用于获取切片 `part` 的长度。由于截取范围是 `[1:3]`，包含索引 `1` 和 `2` 两个元素，因此长度为 `3 - 1 = 2`，输出结果为 `2`。   

part只截取scores俩部分  part[0] 和scores [1]实际上指的是同一个底层数组



##### 切片的大部分入门问题

切片的大部分入门问题都可以从起点，长度，容量和是否共享底层元素四件事拆开分析。需要独立副本时用明确的新空间和复制，需要追加时接住append的返回值。





# 映射与字符串

### Map（字典/映射）

~~~go
func main() {
    scores := map[string]int{
        "小林": 85,
        "小陈": 92
    }
    fmt.Println(scores["小林"])
    fmt.Println(scores["小王"])
    fmt.Println(len(scores))     //查找map的长度
}

//输出结果
85
0
2
~~~

这段代码展示了 Go 语言中 **Map（字典/映射）** 的基本定义、键值查找机制以及长度获取方式：   

1. **Map 的初始化**

   - `scores := map[string]int{"小林": 85, "小陈": 92}` 创建并初始化了一个键（Key）类型为 `string`、值（Value）类型为 `int` 的哈希表。此时，map 中包含了 `"小林": 85` 和 `"小陈": 92` 两组键值对。

   ~~~go
   m := map[KeyType]ValueType{
       key1: value1,
       key2: value2,
   }
   ~~~

      

2. **存在的键查找**

   - `fmt.Println(scores["小林"])` 通过键 `"小林"` 去 map 中查找对应的值。因为该键存在，所以成功输出了其对应的值 **`85`**。   

3. **不存在的键查找（零值规则）**

   - `fmt.Println(scores["小王"])` 尝试查找键 `"小王"`。   
   - **核心规则**：在 Go 语言中，如果通过键读取 map 中**不存在的数据**，程序不会引发运行时错误（Panic），而是会安全地返回该值类型的**零值（Zero Value）**。由于值的类型是 `int`，其零值是 `0`，因此这里会打印出 **`0`**。   
   - *提示*：在实际开发中，如果想区分一个键是“真的存有 0 分”还是“根本不存在”，通常会使用 Go 的“逗号 ok”惯用法来判断：`val, ok := scores["小王"]`。

   ~~~go
   package main
   
   import "fmt"
   
   func main() {
       // 初始化 map，假设“小王”这次考试考了 0 分
       scores := map[string]int{
           "小王": 0,
           "小林": 85,
       }
   
       // 场景 1：查询“小王”（存在，且分数为 0）
       val, ok := scores["小王"]
       if ok {
           fmt.Printf("“小王”存在于成绩单中，分数是：%d 分\n", val)
       } else {
           fmt.Println("“小王”根本不存在于成绩单中")
       }
   
       // 场景 2：查询“小陈”（根本不存在于 map 中）
       val, ok = scores["小陈"]
       if ok {
           fmt.Printf("“小陈”存在于成绩单中，分数是：%d 分\n", val)
       } else {
           fmt.Println("“小陈”根本不存在于成绩单中")
       }
   }
   ~~~

   

4. **获取 Map 的长度**

   - `fmt.Println(len(scores))` 使用内置的 `len()` 函数来获取当前 map 中存储了多少对键值。因为里面有两对数据（小林和小陈），所以输出结果为 **`2`**。   



#####  Map的“逗号 ok”惯用法以及内置的 **`delete` 删除函数**机制：   

~~~go
func main() {
    scores := map[string]int{"小林": 0}
    score, ok := scores["小林"]
    fmt.Println(score, ok)
    delete(scores, "小林")
    score, ok = scores["小林"]
    fmt.Println(score, ok)
}

//输出结果
0 true
0 false
~~~

这段代码主要演示了 Go 语言中 Map（字典）的“逗号 ok”惯用法以及内置的 **`delete` 删除函数**机制：   

1. **Map 初始化与初次查询**
   - `scores := map[string]int{"小林": 0}` 创建了一个 map，里面包含键 `"小林"`，其对应的值是 `0`。   
   - `score, ok := scores["小林"]` 使用了“逗号 ok”惯用法来检查键是否存在。因为 `"小林"` 确实存在于 map 中，所以 `score` 得到 `0`，`ok` 得到 `true`。   
   - `fmt.Println(score, ok)` 打印输出：`0 true`。   
2. **`delete` 函数的机制**
   - `delete(scores, "小林")` 是 Go 语言内置的用于删除 map 中键值对的函数。   
   - **机制**：它接受两个参数（目标 map 和要删除的键）。如果该键存在，它会直接将该键值对从 map 中移除；如果该键本来就不存在，`delete` 会静默处理，不会引发任何错误（是安全的）。   
3. **删除后的再次查询**
   - `score, ok = scores["小林"]` 在执行完删除操作后，再次尝试查找 `"小林"`。   
   - **机制**：由于该键已经被 `delete` 彻底删除了，此时它已不存在于 map 中。因此，`score` 返回了 `int` 类型的零值 `0`，而 `ok` 的值变成了 `false`（用来明确告诉你：这个键已经不在 map 里了）。   
   - `fmt.Println(score, ok)` 打印输出：`0 false`。   



### **Nil Map（未初始化的字典）** 的特性

只声明一个map变量但是没有创建实际映射时，他的零值是nil

~~~go
func main() {
    var scores map[string]int
    fmt.Println(scores == nil)
    fmt.Println(scores["小林"])
    scores = make(map[string]int)
    scores["小林"] = 85
    fmt.Println(scores["小林"])
}

//输出结果
true
0
85
~~~

这段代码主要演示了 Go 语言中 **Nil Map（未初始化的字典）** 的特性、安全读取规则以及通过 `make` 进行初始化的机制：   

1. **声明未初始化的 Map (`nil map`)**
   - `var scores map[string]int` 仅声明了一个 map 变量，但没有进行初始化。在 Go 语言中，未初始化的 map 其默认值就是 `nil`。   
   - `fmt.Println(scores == nil)` 因此会输出 `true`。   
2. **从 Nil Map 中读取数据（安全机制）只能读不能写入**
   - `fmt.Println(scores["小林"])` 尝试从一个 `nil` map 中读取数据。   
   - **机制**：在 Go 语言中，**从一个 `nil` map 中读取数据是完全安全的**，程序不会抛出异常（Panic），而是会直接返回该值类型的零值（`int` 的零值是 `0`），所以这里会打印出 `0`。   
   - *注意*：虽然读取是安全的，但如果尝试向一个 `nil` map 中**写入**数据（例如执行 `scores["小林"] = 85`），会导致运行时恐慌（Panic）。
3. **使用 `make` 函数进行初始化**
   - `scores = make(map[string]int)` 使用 `make` 函数为 map 分配底层的数据结构内存。经过这一步后，`scores` 不再是 `nil`，变成了一个可以正常读写的有效 map。   
4. **写入数据与再次读取**
   - `scores["小林"] = 85` 成功向初始化后的 map 中写入了键值对。   
   - 最后通过 `fmt.Println(scores["小林"])` 再次读取，此时成功获取到了刚存入的数据，输出结果为 `85`。   



### 字符串（String）底层编码原理与遍历机制

~~~go
func main() {
    text := "Go语"
    fmt.Println(len(text), utf8.RuneCountInString(text))
    for position, ch := range text {
        fmt.Printf("%d %c\n", position, ch)
    }
}

//输出结果
5 3
0 G
1 o
2 语
//第一个字符 G 的索引是 0。
//第二个字符 o 的索引是 1。
//第三个字符 语 的索引直接跳到了 2。
//为什么？ 因为前两个英文字符各占 1 个字节（分别占据索引 0 和 1），而中文字符 语 占了 3 个字节（占据索引 2、3、4）。所以下一个字符如果存在，它的起始索引就会变成 5。
~~~

 1.字节长度 vs 字符数量

- **`len(text)`**：获取的是字符串占用的**字节数（Bytes）**。
  - 在 Go 语言中，字符串底层是由字节组成的。**英文字母 `G` 和 `o` 各占 1 个字节，而中文字符 `语` 在 UTF-8 编码中占 3 个字节。**因此，`len("Go语")` 的结果是 `1 + 1 + 3 = 5`。
- **`utf8.RuneCountInString(text)`**：获取的是字符串的**实际字符数量（Unicode 字符 / Rune）**。
  - 字符串中实际包含三个字符：`G`、`o`、`语`，因此结果是 `3`。

2. `for range` 遍历字符串的奇妙机制

当使用 `for position, ch := range text` 遍历字符串时，Go 语言表现得非常聪明：它不是按字节死板地切，而是**自动按 Unicode 字符（Rune）进行解码**：

- **`position`**：返回的是当前字符在底层字节数组中的**起始字节索引（Byte Index）**。
- **`ch`**：返回的是当前的 **Rune 字符**（使用 `%c` 占位符可以将其打印成人类可读的文字）。

### 高频核心字符串处理函数（**`strings` 包** ）

~~~go
// 高频核心字符串处理函数，需要strings包的支持
func main() {
	text := " Go基础 "
	clean := strings.TrimSpace(text)
	fmt.Println(clean)
	fmt.Println(strings.Contains(clean, "Go"))
	parts := strings.Split("小林,85", ",")
	fmt.Println(parts[0], parts[1])
}
~~~

1. **去除首尾空白字符 (`strings.TrimSpace`)**
   - **代码**：`clean := strings.TrimSpace(text)`   
   - **解释**：定义了一个带前后空格的字符串 `" Go基础 "`，通过 `strings.TrimSpace` 函数过滤掉了开头和结尾的空格、制表符等空白，返回处理后的新字符串 `"Go基础"`，并将其打印出来。   
2. **检查子串是否存在 (`strings.Contains`)**
   - **代码**：`fmt.Println(strings.Contains(clean, "Go"))`   
   - **解释**：检查 `clean` 变量中是否包含子串 `"Go"`。因为字符串中确实含有 `"Go"`，所以该函数返回布尔值 `true` 并打印。   
3. **按分隔符切割字符串 (`strings.Split`)**
   - **代码**：`parts := strings.Split("小林,85", ",")`   
   - **解释**：使用逗号 `,` 作为分隔符，将 `"小林,85"` 切割成多个片段。该函数会返回一个字符串切片（`[]string`），其内容为 `["小林", "85"]`。 本且俩个都是string类型，85不是int类型。  
4. **访问切片元素**
   - **代码**：`fmt.Println(parts[0], parts[1])`   
   - **解释**：通过索引 `0` 和 `1` 分别取出切片中的第一个元素（`"小林"`）和第二个元素（`"85"`）并打印输出。   

 核心机制与原理

这段代码主要展示了 Go 语言标准库中 **`strings` 包** 的常用文本处理能力。其背后的设计与底层原理体现在以下几个方面：

1. **字符串的不可变性（Immutability）**
   - 在 Go 语言中，字符串是只读的底层字节数组（Byte Slices）。
   - 像 `strings.TrimSpace` 这样看似“修改”字符串的函数，实际上并不会去改动原内存中的字符串，而是会**创建并返回一个新的字符串**。
2. **高效的切片返回值（Slice Allocation）**
   - `strings.Split` 返回的 `parts` 是一个切片（`[]string`）。Go 的切片底层是一个包含“指向底层数组的指针、长度、容量”的小结构体。
   - 切割操作不会对原字符串的每个字符进行深拷贝，而是高效地通过指针和长度来标记切片出来的各个子串位置。
3. **标准库导入约束**
   - 要使用这些字符串函数，必须在源码文件的头部正确导入标准库：**strings**

### 逗号 ok”惯用法

~~~go
func main() {
    scores := map[string]int{"小林": 0}
    value, ok := scores["小王"]
    fmt.Println(value, ok)
    value, ok = scores["小林"]
    fmt.Println(value, ok)
}

//输出结果
0 false
0 true
~~~

**为什么需要 `ok`？**

- 在很多语言中，如果去 Map 里找一个不存在的键，可能会返回特定的空值（如 Python 的 `None`）。但在 Go 强类型语言中，如果 Map 的值是 `int`，找不到时必须返回一个具体的 `int`（即零值 `0`）。
- 这就带来了管理上的歧义：**返回的 `0`，究竟是因为键不存在得到的默认零值，还是因为这个键真的存了一个数值 `0`？**
- 为了彻底解决这个歧义，Go 设计了“逗号 ok”机制。

**底层设计原理**

- 当执行 `val, ok := myMap[key]` 时，Go 的 Map 底层查找会返回两个独立的指标：
  - 第一个指标是**对应的值**（如果没找到，就返回类型零值）。
  - 第二个指标是一个布尔值（`ok`），它直接反映该键在底层哈希表中**是否真实存在**。
- 这种设计既保证了类型的安全性，又让开发者能够100%准确地判断键的存亡状态。

---

# 函数与延迟执行

### 泛型函数

~~~go
func 函数名[T 约束类型](参数名 参数类型) 返回值类型 {
    // 函数体逻辑
}
~~~

~~~go
func greet(name string) string {
	return "Hello," + name
}

func main() {
	a := greet("xiaowang")
	println(a)
}
~~~

### 函数一次交回多个结果

~~~go
func divide(total, count int) (int, bool) {
    if count == 0 { return 0, false }
    return total / count, true
}
func main() {
    result, ok := divide(10, 2)
    fmt.Println(result, ok)
}

//输出结果
5 true
~~~

 逐行代码解释

1. **定义带多个返回值的函数：`divide`**
   - **代码**：`func divide(total, count int) (int, bool)`   
   - **解释**：定义了一个名为 `divide` 的函数，接收两个整型参数 `total`（被除数）和 `count`（除数）。它的特别之处在于**返回两个值**：一个计算结果（`int`）和一个状态标志（`bool`）。   
2. **分母为零的安全防护**
   - **代码**：`if count == 0 { return 0, false }`   
   - **解释**：在计算机中，整数除以 0 会引发严重错误。这里做了前置防御：如果除数 `count` 等于 0，直接返回零值 `0` 和代表失败的 `false`，从而避免程序崩溃。   
3. **正常的除法计算**
   - **代码**：`return total / count, true`   
   - **解释**：如果除数安全，执行正常的除法运算 `total / count`，并与代表成功的 `true` 一起返回。   
4. **主函数中调用并接收多个返回值**
   - **代码**：`result, ok := divide(10, 2)`   
   - **解释**：调用 `divide` 函数计算 `10 / 2`，并将返回的计算结果存入 `result`，将状态布尔值存入 `ok`。   
5. **打印输出**
   - **代码**：`fmt.Println(result, ok)`   
   - **解释**：将结果和状态打印到控制台，输出结果为 **`5 true`**。   

核心机制与原理：函数的“多返回值”特性

这段代码完美展示了 Go 语言在处理**可能出错或带有状态的操作**时最经典的编程范式：

1. **原生支持多返回值**
   - 许多传统编程语言（如 Java、C++）的函数通常只能返回一个值。如果想返回附加状态，往往需要包装成一个对象。
   - 而 Go 语言原生在语法层面支持**多个返回值**（在签名中用括号括起类型），这使得编写工具函数时非常轻量、优雅。
2. **与 Map “逗号 ok”惯用法的异曲同工之妙**
   - 还记得前面学过的 Map 查找吗？（`value, ok := scores["小王"]`）。   
   - 这里的 `divide` 函数设计完全遵循了同样的哲学：**不仅把数据交给你，还把“这个数据是否有效/安全”的状态（`ok`）诚实地同步告诉你**。
   - 这种设计彻底避免了使用特殊错误码（比如用 `-1` 代表失败，但 `-1` 也有可能是正常的计算结果）带来的混淆，让代码的逻辑极其健壮和安全！



### 可变数量的参数和切片展开符（...）

~~~go
func sum(values ...int) int {
    total := 0
    for _, value := range values { total += value }
    return total
}
func main() {
    scores := []int{80, 90}
    fmt.Println(sum(1, 2, 3))
    fmt.Println(sum(scores...))
}

//输出结果
6
170
~~~

1. **定义可变参数函数：`sum`**
   - **代码**：`func sum(values ...int) int`   
   - **解释**：定义了一个名为 `sum` 的函数，它可以接收**零个或多个** `int` 类型的参数。这里的 `...int` 表示这是一个可变参数。   
2. **初始化累加器并遍历求和**
   - **代码**：`total := 0` 随后通过 `for _, value := range values { total += value }`   
   - **解释**：初始化一个变量 `total` 用来记录总和。使用 `for range` 循环遍历传入的所有数值（其中索引被下划线 `_` 忽略），把每一个 `value` 累加到 `total` 中。最后返回计算好的 `total`。   
3. **创建切片与直接传参调用**
   - **代码**：`scores := []int{80, 90}`，接着 `fmt.Println(sum(1, 2, 3))`   
   - **解释**：首先定义了一个整型切片 `scores` 包含 `80` 和 `90`。然后第一次调用 `sum(1, 2, 3)`，直接传入三个独立的整数，输出结果为 `6`。   
4. **切片解包（Spread）传参**
   - **代码**：`fmt.Println(sum(scores...))`   
   - **解释**：第二次调用时，传入了 `scores...`。这里利用了 Go 语言的切片展开特性，把切片里的元素打散后逐个传给可变参数函数，输出结果为 `170`（`80 + 90`）。   

三、 核心机制与原理

这段代码展示了 Go 语言在处理**不固定数量参数**时的两个核心特性：

1. **可变参数（Variadic Parameters）**
   - 语法上使用 `...类型`（如 `...int`）。在函数内部，`values` 的本质会被当作一个切片（`[]int`）来处理。
   - 这极大地简化了像求和、拼接这类需要处理动态个数参数的函数编写。
2. **切片展开操作符（Spread Operator `...`）**
   - 当你手中已经有一个切片（如 `scores`），而某个函数需要接收独立的可变参数时，不能直接把切片传进去（类型不匹配）。
   - 在切片后面加上三个点 `...`（即 `scores...`），就能告诉编译器：“**请把这个切片里的所有元素解开，像一个个独立参数那样传递过去**”。这个设计让切片和可变参数函数之间能够无缝衔接。



### defer在函数退出时执行（后进先出）

~~~go
func main() {
    defer fmt.Println("先登记")
    defer fmt.Println("后登记")
    fmt.Println("处理中")
}

//输出结果
处理中
后登记
先登记
~~~

1. **第一次使用 `defer` 注册延迟任务**
   - **代码**：`defer fmt.Println("先登记")`   
   - **解释**：`defer` 是 Go 语言特有的关键字，用来修饰一个函数调用。被 `defer` 修饰的函数不会立即执行，而是会被“挂起”并推迟到**当前函数即将返回（退出）的前一刻**才执行。这里注册了要打印的 `"先登记"`。   
2. **第二次使用 `defer` 注册延迟任务**
   - **代码**：`defer fmt.Println("后登记")`   
   - **解释**：同样使用 `defer` 注册了第二个延迟执行的任务，打印 `"后登记"`。   
3. **立即执行的普通语句**
   - **代码**：`fmt.Println("处理中")`   
   - **解释**：这是一条普通的打印语句。没有被 `defer` 修饰的代码会按照书写顺序**正常立即执行**。   

 核心机制与原理：`defer` 的“后进先出”（LIFO）栈结构

1. **执行流程推演**：
   - **第一步**：程序进入 `main` 函数，遇到第一行 `defer`，把 `"先登记"` 压入延迟执行的栈中。
   - **第二步**：遇到第二行 `defer`，把 `"后登记"` 压入栈中（此时它在栈顶）。
   - **第三步**：执行到 `fmt.Println("处理中")`，**立即在控制台打印出 `"处理中"`**。
   - **第四步**：`main` 函数即将结束，开始逆序弹出并执行 `defer` 队列中的任务：
     - 栈顶的 `"后登记"` 先被弹出执行。
     - 压底的 `"先登记"` 后被弹出执行。
2. **为什么要设计成“后进先出”（LIFO）？**
   - 在实际开发中，`defer` 最常用来做**资源清理**（例如关闭文件、关闭数据库连接、释放锁）。
   - 想象一下，如果你先打开了文件 A，后打开了文件 B，在关闭时，通常需要**先关闭后打开的资源（B）**，再关闭先打开的资源（A），以防出现依赖死锁或资源未释放。Go 语言的 `defer` 这种倒序执行的设计，完美契合了这种资源释放的逻辑！

### defer延迟调用是保存当时的值

~~~go
func main() {
    a := 80
    defer fmt.Println(a)
    a := 90
    fmt.Println(a)
}

//输出结果
90
80
~~~

- 先执行普通输出
- defer调用的参数在登记时准备
- defer延迟调用是保存当时的值



---

# 指针和结构体

### 指针（Pointers）与内存操控

~~~go
func main() {
    score := 80
    pointer := &score
    *pointer = 85
    fmt.Println(score)
    fmt.Println(*pointer)
}
~~~

1. **获取变量的内存地址（取址操作）**
   - **代码**：`pointer := &score`   
   - **解释**：使用取址操作符 `&` 获取了变量 `score` 在计算机内存中的实际存储地址，并将其赋值给变量 `pointer`。此时 `pointer` 的类型是一个整型指针（`*int`）。   
2. **通过指针修改内存中的值（解引用赋值）**
   - **代码**：`*pointer = 85`   
   - **解释**：使用解引用操作符 `*` 顺着指针找到它所指向的内存地址，并把该内存位置上的数据直接修改为 `85`。因为这块内存空间正是 `score` 所在的地址，所以变量 `score` 的值也随之改变了。   
3. **打印变量与指针的值**
   - **代码**：`fmt.Println(score)` 和 `fmt.Println(*pointer)`   
   - **解释**：分别直接打印变量 `score`，以及通过 `*pointer` 解引用获取指针指向的值。两者的控制台输出结果均为 **`85`** 和 **`85`**。   

**核心机制与原理：指针（Pointers）与内存操控**

1. **什么是内存地址？**
   - 变量在计算机底层占用一块内存空间，每个空间都有一个唯一的编号（即内存地址）。
2. **取址符（`&`）**
   - 放在变量前面（如 `&score`），作用是“拿到这个变量在内存中的家门牌号（地址）”。   
3. **解引用符（`\*`）**
   - 放在指针变量前面（如 `*pointer`），作用是“顺着门牌号去找房间，并对房间里面的东西进行读取或修改”。   
4. **核心意义**
   - **通过指针，我们可以在不同的函数或代码段中直接操作同一块内存数据，而不需要通过复制变量副本传递，这在提高程序性能和实现“引用传递”时至关重要。**



### Go 语言的 值传递（Pass by Value）

~~~go
func addLocal(score int) { score += 5 }
func addOriginal(score *int) { *score += 5 }
func main() {
    score := 80
    addLocal(score)
    fmt.Println(score)
    addOriginal(&score)
    fmt.Println(score)
}
~~~

1. **定义值传递函数：`addLocal`**
   - **代码**：`func addLocal(score int) { score += 5 }`   
   - **解释**：接收一个整型参数 `score`。在函数内部，对这个参数加了 5，但由于这是按值传递，修改的只是函数内部的局部副本。   
2. **定义指针传递函数：`addOriginal`**
   - **代码**：`func addOriginal(score *int) { *score += 5 }`   
   - **解释**：接收一个整型指针 `*int` 作为参数。通过解引用操作 `*score`，直接去修改该指针所指向的内存地址上的数值。   
3. **初始化变量并测试值传递**
   - **代码**：`score := 80` 随后调用 `addLocal(score)`   
   - **解释**：定义初始分数为 `80`。调用 `addLocal(score)` 时，传入的是 `score` 的复本，函数内部的加 5 无法影响外面的变量。   
   - **输出**：紧接着执行 `fmt.Println(score)`，打印出来的结果仍然是 **`80`**。   
4. **测试指针传递并打印最终结果**
   - **代码**：`addOriginal(&score)` 随后 `fmt.Println(score)`   
   - **解释**：调用 `addOriginal(&score)`，把 `score` 的内存地址（`&score`）传了进去。函数直接修改了原内存中的数据，使得变量本身发生了改变。   
   - **输出**：最后再次打印 `score`，结果变成了 **`85`**。   

**核心机制与原理：Go 语言的“值传递（Pass by Value）”**

这段代码极其经典地对比了 Go 语言中**值传递**与**指针传址**的区别：

1. **Go 语言只有“值传递”**
   - 无论是基础类型（如 `int`、`string`）还是结构体、切片，当它们作为参数传给函数时，Go 永远会把数据 **复制一份** 传递过去。
   - 这就是为什么 `addLocal(score)` 运行后，外面的 `score` 毫无变化——因为它改的只是复制品。
2. **如何实现“修改原变量”？**
   - 如果想在一个函数内部改变外部变量的值，必须传递该变量的**内存地址（指针，如 `\*int`）**。   
   - 函数拿到地址后，通过解引用（`*score`）直接操作原变量在内存中的实际存储空间，从而实现状态的改变。   



### 结构体（Struct）

在 Go 语言中，**结构体（Struct）** 是一种自定义的**复合数据类型**。它可以把多个不同类型的数据项（称为“字段”/ Field）打包组合在一起，形成一个新的逻辑实体。   

~~~go
type Student struct {
    Name  string
    Score int
}
func main() {
    student := Student{Name: "小林", Score: 80}
    fmt.Println(student.Name, student.Score)
    fmt.Println(student)
}

//输出结果
小林 80
{小林 80}
~~~

1. **定义结构体类型 (`struct`)**
   - **代码**：`type Student struct { Name string; Score int }`   
   - **解释**：使用 `type` 和 `struct` 关键字定义了一个自定义的复合数据类型叫 `Student`。它内部包含两个字段（Field）：一个是字符串类型的 `Name`，另一个是整型类型的 `Score`。   
2. **结构体实例化与初始化**
   - **代码**：`student := Student{Name: "小林", Score: 80}`   
   - **解释**：在 `main` 函数中，通过指定字段名称的方式（键值对初始化）创建了一个 `Student` 结构体的实例，赋值给变量 `student`。此时，`Name` 被赋为 `"小林"`，`Score` 被赋为 `80`。   
3. **通过点操作符访问字段**
   - **代码**：`fmt.Println(student.Name, student.Score)`   
   - **解释**：使用点（`.`）操作符分别访问结构体实例内部的 `Name` 和 `Score` 字段，并将其打印到控制台。   
   - **控制台输出结果**：`小林 80`。   

**这段代码展示了 Go 语言中处理复杂结构时最核心的数据组织形式：**

1. **告别传统“类（Class）”，拥抱结构体**
   - Go 语言的设计哲学中没有面向对象语言（如 Java、C++）中经典的“类（Class）”和“继承”概念。
   - 取而代之的是 **结构体（Struct）**。它纯粹用于把相关联的多个不同类型的数据组合封装在一起，体现了组合优于继承的设计思想。
2. **结构体是“值类型（Value Type）”**
   - 在 Go 语言中，结构体属于值类型。**这意味着当把一个结构体变量赋值给另一个变量，或者将它作为参数传递给函数时，整个结构体的数据会被完整复制一份，而不是传递内存地址。**
   - 这一特性保证了数据的安全性，但如果结构体非常庞大，通常会配合指针（如 `*Student`）来使用，以避免复制带来的性能开销。



### 结构体赋值也是值复制

~~~go
type Student struct {
    Name  string
    Score int
}
func main() {
    first := Student{Name: "小林", Score: 80}
    second := first
    second.Score = 90
    fmt.Println(first.Score, second.Score)
}

//输出
80 90
~~~

1. **定义结构体类型**
   - **代码**：`type Student struct { Name string; Score int }`   
   - **解释**：定义了一个名为 `Student` 的结构体，包含两个字段：`Name`（字符串类型）和 `Score`（整型）。   
2. **实例化结构体并初始化**
   - **代码**：`first := Student{Name: "小林", Score: 80}`   
   - **解释**：创建了一个 `Student` 实例 `first`，其名字为 `"小林"`，分数为 `80`。   
3. **赋值与复制**
   - **代码**：`second := first`   
   - **解释**：将 `first` 赋值给新变量 `second`。**注意**：这里发生了一次结构体的**完整数值复制**，`second` 成为了一个独立的、拥有自己独立内存空间的全新对象。   
4. **修改副本的字段值**
   - **代码**：`second.Score = 90`   
   - **解释**：将 `second` 的 `Score` 字段修改为 `90`。由于它和 `first` 是互不干扰的独立副本，这次修改**不会**影响到 `first`。   
5. **打印输出**
   - **代码**：`fmt.Println(first.Score, second.Score)`   
   - **解释**：分别打印出两个结构体实例的分数，结果分别为 `80` 和 `90`。   

**核心机制与原理**：结构体的“值类型”特性

这段代码完美体现了 Go 语言中结构体的一大核心底层机制：

1. **结构体是值类型（Value Type）**
   - 在 Go 语言中，结构体和基础数据类型（如 `int`、`float`、`string`）一样，都属于**值类型**。
   - 当你把一个结构体变量赋值给另一个变量（`second := first`），或者把结构体传给函数时，Go 编译器会**把整块内存中的所有数据完整复制一份**给对方。   
2. **相互独立，互不影响**
   - 因为是复制了一份独立的副本，所以修改 `second` 的内部属性（如 `second.Score = 90`）只会改变它自己，原对象 `first` 的数据会完好无损地保留下来（保持为 `80`）。   
   - **对比提示**：如果你希望多个变量指向并共享同一块结构体数据（修改一个，另一个也跟着变），在 Go 中就需要显式使用**结构体指针**（如 `*Student`）。



### *通过结构体指针修改原数据

这段代码展示了 Go 语言中如何通过**指针**来突破“值传递”的限制，从而修改原结构体数据：

重点：（要想在函数内部**改变外部结构体的状态**，必须传递它的**内存地址（指针）** ）

~~~go
type Student struct { Score int }
func setScore(student *Student) {
    student.Score = 85   //go会自动解引用，所以写成(*student).Score = 85
}
func main() {
    s := Student{Score: 80}
    setScore(&s)     //要想在函数内部改变外部结构体的状态，必须传递它的内存地址（指针）。  
    fmt.Println(s.Score)
}

//输出结果
85
~~~

1. **定义结构体**
   - **代码**：`type Student struct { Score int }`   
   - **解释**：定义了一个名为 `Student` 的结构体，内部只有一个整型字段 `Score`。   
2. **定义接收指针的函数**
   - **代码**：`func setScore(student *Student) { student.Score = 85 }`   
   - **解释**：定义了一个名为 `setScore` 的函数，它接收的参数是一个结构体 `Student` 的**指针**（类型为 `*Student`）。在函数内部，直接将该指针指向的对象的 `Score` 字段修改为 `85`。   
3. **初始化结构体与传址调用**
   - **代码**：`s := Student{Score: 80}` 随后 `setScore(&s)`   
   - **解释**：首先创建了一个分数为 `80` 的学生实例 `s`。接着，使用取址符 `&s` 获取了 `s` 的内存地址，并将其传给 `setScore` 函数。   
4. **打印输出**
   - **代码**：`fmt.Println(s.Score)`   
   - **解释**：打印变量 `s` 的分数，由于函数内部通过指针修改了原件，输出结果变成了 **`85`**。   

核心机制与原理：结构体与指针的结合

**这段代码展示了 Go 语言中如何通过指针来突破“值传递”的限制，从而修改原结构体数据：**

1. **结构体是值类型**
   - 我们前面学过，结构体在默认情况下是值类型。如果直接传 `s`（而不是 `&s`），函数内部拿到的只是一个复制品，修改它对外部的 `s` 毫无影响。
   - 要想在函数内部**改变外部结构体的状态**，必须传递它的**内存地址（指针）**。   
2. **Go 语言的语法糖：自动解引用（Automatic Dereferencing）**
   - 你可能会好奇：在 `setScore` 函数里，`student` 是一个指针（`*Student`），为什么可以直接写 `student.Score = 85`，而不需要写繁琐的 `(*student).Score = 85` 呢？
   - **机制原理**：Go 语言编译器非常智能。当你通过结构体指针去访问或修改字段时，如果发现左边是个指针，Go 会**自动帮你做解引用**，让你可以像操作普通结构体一样流畅地编写代码。



### *克隆体与本体

~~~go
type Student struct { Score int }
func main() {
    first := Student{Score: 80}
    second := first    //复制
    pointer := &first   //取地址
    second.Score = 90
    pointer.Score = 85
    fmt.Println(first.Score, second.Score)
}

//输出结果
85 90
~~~

1. **定义结构体**
   - **代码**：`type Student struct { Score int }`   
   - **解释**：定义了一个名为 `Student` 的结构体，内部包含一个整型字段 `Score`。   
2. **初始化实例**
   - **代码**：`first := Student{Score: 80}`   
   - **解释**：创建了一个 `Student` 实例 `first`，其分数为 `80`。   
3. **结构体赋值（值拷贝）**
   - **代码**：`second := first`   
   - **解释**：将 `first` 赋值给 `second`。由于结构体是值类型，这里会把 `first` 的数据**完整复制一份**给 `second`，两者拥有各自独立的内存空间。   
4. **获取指针（取址）**
   - **代码**：`pointer := &first`   
   - **解释**：使用取址符 `&` 获取了原实例 `first` 在内存中的实际地址，并赋值给指针变量 `pointer`。   
5. **修改副本的值**
   - **代码**：`second.Score = 90`   
   - **解释**：修改副本 `second` 的分数。因为它是独立复制出来的，这次修改**不会**影响到 `first`。   
6. **通过指针修改原实例的值**
   - **代码**：`pointer.Score = 85`   
   - **解释**：通过指针 `pointer` 直接去修改原实例 `first` 的分数。此时，`first` 的分数被成功更新为 `85`。   
7. **打印输出**
   - **代码**：`fmt.Println(first.Score, second.Score)`   
   - **解释**：分别打印 `first` 和 `second` 的分数，输出结果为 **`85 90`**。   

这段代码完美地把前面学过的两个核心机制结合在了一起：

1. **结构体的“值类型”独立复制**
   - 当执行 `second := first` 时，Go 进行了内存层面的值拷贝，使 `second` 成为了一个独立的个体，修改它不会波及 `first`。   
2. **结构体指针与“自动解引用”语法糖**
   - 当执行 `pointer.Score = 85` 时，`pointer` 是一个结构体指针（`*Student`）。   
   - 按照严谨的底层逻辑，修改指针指向的字段应该写成 `(*pointer).Score = 85`。但 Go 语言内置了**自动解引用机制**，允许你直接用点操作符（`.`）通过指针去访问或修改字段，让代码写起来更加优雅、自然。



# 方法与接口

- **结构体（Struct）** 决定了“我是谁”（数据的载体）。
- **方法（Method）** 决定了“我能做什么”（行为的绑定）。
- **接口（Interface）** 决定了“我们有哪些共同的标准”（多态与解耦的契约）。

通过这三者的组合，Go 语言在没有传统面向对象“类和继承”的情况下，构建出了极其灵活、高内聚且低耦合的代码世界。



### Go 语言的方法（Method）与值接收者

~~~go
type Student struct { Score int }
func (s Student) Passed() bool {
    return s.Score >= 60
}
func main() {
    student := Student{Score: 80}
    fmt.Println(student.Passed())
}

//输出结果
true
~~~

1. **定义结构体**
   - **代码**：`type Student struct { Score int }`   
   - **解释**：定义了一个名为 `Student` 的结构体，内部包含一个整型字段 `Score`。   
2. **定义结构体方法**
   - **代码**：`func (s Student) Passed() bool { return s.Score >= 60 }`   
   - **解释**：为 `Student` 结构体绑定了一个名为 `Passed` 的方法。该方法返回一个布尔值（`bool`）。内部通过 `s.Score >= 60` 判断分数是否及格。   
3. **实例化结构体并调用方法**
   - **代码**：`student := Student{Score: 80}` 随后 `fmt.Println(student.Passed())`   
   - **解释**：在 `main` 函数中创建了一个 `Student` 实例，分数为 `80`。接着通过点操作符（`.`）调用了它的 `Passed()` 方法，并把结果打印到控制台。   

 核心机制与原理：Go 语言的方法（Method）与值接收者

这段代码展示了 Go 语言中面向对象风格的函数绑定机制：

1. **什么是方法（Method）？**
   - Go 语言虽然没有传统面向对象语言（如 Java 或 C++）中的 `class` 关键字，但它允许为自定义类型（如结构体）定义**方法**。
   - 方法与普通函数的区别在于，它多了一个接收者（Receiver）参数——即函数名前面的 `(s Student)`。这让 `Passed` 方法从属于 `Student`，只有 `Student` 的实例才能调用它。   
2. **值接收者（Value Receiver）机制**
   - 这里的接收者类型是 `Student`（而不是指针 `*Student`），这被称为**值接收者**。
   - 当通过 `student.Passed()` 调用时，Go 语言会把 `student` 实例的值**复制一份**传递给方法内部的 `s`。
   - **适用场景**：因为该方法只需要**读取**数据（判断 `s.Score >= 60`），而不需要修改原结构体的内容，使用值接收者非常安全且直观。   

