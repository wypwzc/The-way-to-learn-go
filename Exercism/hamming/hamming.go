package hamming

import "errors"

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		//errors.New（返回错误信息）
		return 0, errors.New("错误：必须传入长度相同的两个字符串")
	}
	//初始化计数器
	distance := 0
	//遍历字符串的每一个索引位置
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}

	return distance, nil
}

/*解释
这段代码是用 Go 语言编写的，用于计算两个字符串之间的 汉明距离（Hamming Distance）。

汉明距离是一个概念，表示两个相同长度的字符串对应位置上不同字符的数量。简单来说，它衡量了把一个字符串变成另一个字符串需要替换多少个字符。

代码逐行解析
1. 包声明与导入
package hamming

import "errors"
package hamming：声明当前代码属于 hamming 包。

import "errors"：导入 Go 标准库的 errors 包，用于创建和返回错误信息。

2. 函数签名
func Distance(a, b string) (int, error)
定义了一个名为 Distance 的公开函数（首字母大写）。

接收两个字符串参数 a 和 b。

返回两个值：int（计算出的汉明距离）和 error（如果出错则返回错误信息，否则为 nil）。

3. 长度检查（前置校验）
if len(a) != len(b) {
    return 0, errors.New("strings must be of equal length")
}
核心逻辑：汉明距离的前提是两个字符串长度必须完全相同。

如果长度不等，函数会立即终止，返回 0 以及一个错误信息 "strings must be of equal length"（字符串必须长度相等）。

4. 初始化计数器
distance := 0
定义一个整型变量 distance 并初始化为 0，用来记录不匹配的字符个数。

5. 循环与逐字符比较
Go
for i := 0; i < len(a); i++ {
    if a[i] != b[i] {
        distance++
    }
}
使用 for 循环遍历字符串的每一个索引位置（从 0 到 len(a) - 1）。

a[i] 和 b[i] 获取的是字符串在索引 i 处的字节值。

如果对应位置的字节不相等（a[i] != b[i]），说明该位置字符不同，distance 自增 1（distance++）。

6. 返回结果
Go
return distance, nil
循环结束后，返回最终计算出的 distance 数值，并将 error 设为 nil（表示运行成功，无错误）。

技术小贴士
字节级比较：代码中使用 a[i] 访问的是底层字节。
对于常规的 ASCII 字符或生物信息学中的 DNA 序列（如 A、C、G、T），这种方式非常高效且完全正确。
但如果处理的是包含多字节的 UTF-8 字符（如中文汉字），直接按字节索引可能会切断字符，通常需要转为 []rune 处理。
对于经典的汉明距离练习题（如 Exercism 上的 Hamming 题目），这种实现是标准且正确的。
*/
