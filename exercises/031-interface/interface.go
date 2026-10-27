//
// 问题：
// 接口类型被定义为一组方法签名。
//
// 接口类型的值可以持有任何实现了这些方法的值。
//
// 当一个类型实现了接口中包含的所有方法时，
// 就称该类型"实现了"那个接口。
//
// Go 中最常见的接口是 io.Reader 和 io.Writer 接口。

// I AM STILL GOING

package main

import "fmt"

func main() {
	const myNumber = 8
	// 你的任务是实现 Decimal 类型并实现 Number 接口，
	// 这样这些行就不会报错。请注意！
	// 你不需要修改 main() 函数中的其余代码！
	// 一旦当前问题解决，下面这些行取消注释后就应该能够正常运行。

	// d := Decimal(myNumber)
	// PrintNumber(d)

	b := Binary(myNumber)
	PrintNumber(b)

	r := Roman(myNumber)
	PrintNumber(r)
}

type Number interface {
	Numeral() string
}

func PrintNumber(n Number) {
	fmt.Println(n.Numeral())
}

// Binary 和 Roman 类型**隐式地**实现了 Number 接口。
// 与其他语言不同，Go 中没有 "implements" 子句。只需编写
// 方法，编译器就会知道这些方法已被实现。

type Binary int
type Roman int

func (b Binary) Numeral() string {
	return fmt.Sprintf("%b", b)
}

func (r Roman) Numeral() (numeral string) {
	switch r {
	case 0:
		numeral = "0"
	case 1:
		numeral = "I"
	case 2:
		numeral = "II"
	case 3:
		numeral = "III"
	case 4:
		numeral = "IV"
	case 5:
		numeral = "V"
	case 6:
		numeral = "VI"
	case 7:
		numeral = "VII"
	case 8:
		numeral = "VIII"
	case 9:
		numeral = "IX"
	case 10:
		numeral = "X"
	default:
		numeral = "error: roman numeral too large"
	}
	return numeral
}
