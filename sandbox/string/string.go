package main

import (
	"fmt"
	"strings"
)

//字符串（String）底层编码原理与遍历机制,需要unicode/utf8包的支持
/*func main() {
text := "Go语"
fmt.Println(len(text), utf8.RuneCountInString(text))
for position, ch := range text {
	fmt.Printf("%d %c\n", position, ch)
}
*/

// 高频核心字符串处理函数，需要strings包的支持
func main() {
	text := " Go基础 "
	clean := strings.TrimSpace(text)
	fmt.Println(clean)
	fmt.Println(strings.Contains(clean, "Go"))
	parts := strings.Split("小林,85", ",")
	fmt.Println(parts[0], parts[1])
}
