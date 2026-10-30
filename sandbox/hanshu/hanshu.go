package main

func greet(name string) string {
	return "Hello," + name
}

func main() {
	a := greet("xiaowang")
	println(a)
}
