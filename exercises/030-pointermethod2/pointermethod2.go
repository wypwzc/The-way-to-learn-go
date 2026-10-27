// 问题：

// I AM STILL GOING

package main

import "fmt"

func main() {
	db := userDatabase{
		admins: []string{"clair@example.com"},
	}
	// 我们希望能够添加管理员。
	// 创建一个可以向数据库添加用户的方法。
	db.AddAdmin("richie@example.com")
	fmt.Println(db.IsAdmin("richie@example.com"), db.IsAdmin("ned@hack.com"))
}

type userDatabase struct {
	admins []string
}

// 在 Go 中，方法通常位于类型定义下方。

// 如果 email 在管理员数据库中，IsAdmin 返回 true，否则返回 false。
func (u userDatabase) IsAdmin(email string) bool {
	for _, got := range u.admins {
		if got == email {
			return true
		}
	}
	return false
}
