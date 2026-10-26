package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func getNumber(num1 int, num2 int) (int,int) {
	sum:= num1 + num2
	multiply := num1 * num2
	return sum, multiply
}

func main() {
	a := 10
	b := 20

	// sum := add(a, b)
	// additionalSum := add(5, 15)
	// fmt.Println("Sum of a and b is:", sum)
	// fmt.Println("Additional sum is:", additionalSum)
	 s,m := getNumber(a, b)
	fmt.Println("Sum of a and b is:", s)
	fmt.Println("Product of a and b is:", m)
}
