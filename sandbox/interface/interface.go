package main

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
