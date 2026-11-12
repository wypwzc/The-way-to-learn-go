package main

import (
	"Exercism/hamming" // 注意：这里的路径是 "你的模块名/子文件夹名"
	"fmt"
)

func main() {
	// 调用 hamming 包里的 Distance 函数
	dist, err := hamming.Distance("GAGCCTACTAACGGGAT", "ATCGTAATGACGGCCT")

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("计算出的汉明距离是:", dist)
}
