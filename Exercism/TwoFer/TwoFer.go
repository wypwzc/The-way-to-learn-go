package twofer

import "fmt"

// ShareWith 根据传入的名字返回对应的分饼干语句
func ShareWith(name string) string {
	// 如果没有传入名字（即名字为空字符串），则默认使用 "you"
	if name == "" {
		name = "you"
	}

	// 使用 fmt.Sprintf 格式化并返回最终的字符串
	return fmt.Sprintf("One for %s, one for me.", name)
}
