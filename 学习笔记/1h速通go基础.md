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
    scores := map[string]int{"小林": 85, "小陈": 92}
    fmt.Println(scores["小林"])
    fmt.Println(scores["小王"])
    fmt.Println(len(scores))
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

####  字节长度 vs 字符数量

- **`len(text)`**：获取的是字符串占用的**字节数（Bytes）**。
  - 在 Go 语言中，字符串底层是由字节组成的。英文字母 `G` 和 `o` 各占 1 个字节，而中文字符 `语` 在 UTF-8 编码中占 3 个字节。因此，`len("Go语")` 的结果是 `1 + 1 + 3 = 5`。
- **`utf8.RuneCountInString(text)`**：获取的是字符串的**实际字符数量（Unicode 字符 / Rune）**。
  - 字符串中实际包含三个字符：`G`、`o`、`语`，因此结果是 `3`。

#### 2. `for range` 遍历字符串的奇妙机制

当使用 `for position, ch := range text` 遍历字符串时，Go 语言表现得非常聪明：它不是按字节死板地切，而是**自动按 Unicode 字符（Rune）进行解码**：

- **`position`**：返回的是当前字符在底层字节数组中的**起始字节索引（Byte Index）**。
- **`ch`**：返回的是当前的 **Rune 字符**（使用 `%c` 占位符可以将其打印成人类可读的文字）。
