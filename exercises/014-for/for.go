//
// 问题：
// For 语句还可以在子句中有初始化和后置语句。
//
//  for i := 0; i <= 10; i = i + 1 {
//  	fmt.Println(i)
//  }
//
// 上面的 for 循环将从 0 迭代到 10。
// 下面的单条件 for 语句将执行相同的操作：
//
//  i := 0
//  for i <= 10 {
//  	fmt.Println(i)
//  	i = i + 1
//  }
//

// I AM STILL GOING

package main

func main() {
	// 哎呀，这个 for 语法真难记！
	// Gopher 们第一次见到 for 循环语法后，
	// 还需要经常查阅怎么写 for 循环，这是很常见的。
	// 修复这个 for 语句时请记住：
	//  1. 第一个语句是初始化（开始时运行一次）
	//  2. 第二个是条件。条件为 true 时 For 循环继续执行
	//  3. 第三个是后置语句。每次循环结束时运行
	for i <= 3; i := 0;  i = i + 1 {
		fmt.Println(i)
	}
}
