package main

import (
	"fmt"
)

func main() {
	a := 10
	b := &a
	*b = 20
	fmt.Println(a, b, *b)
}
