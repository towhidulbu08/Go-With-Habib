package main

import "fmt"

func add(a, b int) {

	fmt.Println(a + b)
}

func main() {
	add(5, 6) //* standard or named function
	func(a, b int) {
		//* anonymous function
		//* IIFE(immediately invoked function expression)
		c := a + b
		fmt.Println(c)
	}(5, 5)

}

func init() {

	fmt.Println("I'll be called first")
}
