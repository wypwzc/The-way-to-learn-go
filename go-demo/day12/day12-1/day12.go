package main

import "fmt"

func main() {

	scores := []int{80, 90, 70}
	part := scores[1:3]
	fmt.Print(scores)
	fmt.Print(part)
}
