package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	foo()
	for i := 0; i < 10; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	n, err := fmt.Println("Hello, World!", 1234, true)
	fmt.Println(n)
	fmt.Println(err)

	m, _ := fmt.Println("Hello, World!", 1234, true)
	//m, e := fmt.Println("Hello, World!", 1234, true)
	fmt.Println(m)
}

func foo() {
	fmt.Println("I'm in foo")
}
