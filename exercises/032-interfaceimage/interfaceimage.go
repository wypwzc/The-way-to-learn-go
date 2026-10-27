//
// 问题：
// Go 中最有趣的接口之一是 `image.Image` 接口。
//
// 它是一个简单得出奇但范围广泛的接口。只需几行代码，
// 我们就可以定义一整个图像。

// I AM STILL GOING

package main

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"math/rand"
	"os"
)

const imageSize = 100

func main() {
	// 你可能还没有接触过文件操作，这是在 Go 中创建文件的方式。
	// 文件创建可能会失败，因此会返回一个错误。
	fp, err := os.Create("stars.png")
	if err != nil {
		// 如果返回错误，我们调用 log.Fatal 终止程序
		// 以便打印错误信息。
		log.Fatal(err)
	}
	// 现在天空还很暗。让我们用星星填满它。
	// 通过设置 s（stars 类型）的值，我们可以改变
	// 任意像素上出现星星的概率。尝试将其设置为 0 到 100 之间的值
	// 并再次运行代码。
	var s stars
	err = png.Encode(fp, s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("written a starlit sky to stars.png")
}

type stars int

func (s stars) At(i, j int) color.Color {
	// 为了判断像素是否为星星，我们检查
	// 0 到 100 之间的随机整数是否小于 stars 接收者的值。
	if rand.Intn(imageSize*imageSize) < int(s) {
		return color.White
	}
	return color.Black
}

// ColorModel 返回可以从 At 方法中期望的颜色范围。
// 由于我们绘制的是黑白天空，color.GrayModel 对这个用例来说已经足够了。
func (s stars) ColorModel() color.Model {
	return color.GrayModel
}

// Bounds 定义了一个 100x100 的星空网格。
func (s stars) Bounds() image.Rectangle {
	return image.Rect(0, 0, imageSize, imageSize)
}
