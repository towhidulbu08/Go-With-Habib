package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	a := 10
	b := 20

	sum := add(a, b)
	additionalSum := add(5, 15)
	fmt.Println("Sum of a and b is:", sum)
	fmt.Println("Additional sum is:", additionalSum)
}
