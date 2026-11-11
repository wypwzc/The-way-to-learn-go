/*package main

import "fmt"

type Speaker interface {
	Speak() string
}

func MakeSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

type Cat struct{}

func (c Cat) Speak() string {
	return "喵喵喵"
}

type Dog struct{}

func (d Dog) Speak() string {
	return "汪汪汪"
}

func main() {
	c := Cat{}
	d := Dog{}
	MakeSpeak(c)
	MakeSpeak(d)
}
*/ //接口典型例子

package main

import (
	"fmt"
	"strconv"
)

// 定义一个接口
type Good interface {
	settleAccount() int
	orderInfo() string
}

type Phone struct {
	name     string
	quantity int
	price    int
}

func (phone Phone) settleAccount() int {
	return phone.quantity * phone.price
}
func (phone Phone) orderInfo() string {
	return "您要购买" + strconv.Itoa(phone.quantity) + "个" +
		phone.name + "计：" + strconv.Itoa(phone.settleAccount()) + "元"
}

type FreeGift struct {
	name     string
	quantity int
	price    int
}

// 结算账户，计算价格
func (gift FreeGift) settleAccount() int {
	return 0
}
func (gift FreeGift) orderInfo() string {
	return "赠送您" + strconv.Itoa(gift.quantity) + "个" +
		gift.name + "计：" + strconv.Itoa(gift.settleAccount()) + "元"
}

// 计算所有商品的总价
func calculateAllPrice(goods []Good) int {
	var allPrice int
	for _, good := range goods {
		fmt.Println(good.orderInfo())
		allPrice += good.settleAccount()
	}
	return allPrice
}
func main() {
	iPhone := Phone{
		name:     "iPhone",
		quantity: 1,
		price:    8000,
	}
	earphones := FreeGift{
		name:     "耳机",
		quantity: 1,
		price:    200,
	}

	goods := []Good{iPhone, earphones}
	allPrice := calculateAllPrice(goods)
	fmt.Printf("该订单总共需要支付 %d 元", allPrice)
}
