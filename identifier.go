package main

import "fmt"

var hello = 12
var z int

func main() {
	a := 1
	fmt.Println(a)
	x := 42
	fmt.Println(x)
	x = 22
	fmt.Println(x)
	y := 100 + 55
	fmt.Println(y)

	var hi = 43
	fmt.Println(hi)
	test()
}

func test() {
	fmt.Println("This is var", hello)
	fmt.Println("This is var", z)
}
