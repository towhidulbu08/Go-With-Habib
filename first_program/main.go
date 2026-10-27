package main

import "fmt"

func printWelcomeMsg() {

	//Print welcome message

	fmt.Println("Welcome to the application")

}

func getUserName() string {

	//Get user name as input

	var name string

	fmt.Println("Enter your name -")
	fmt.Scanln(&name)
	return name
}

func getTwoNumbers() (int, int) {

	var num1 int
	var num2 int
	fmt.Println("Enter first Number: ")
	fmt.Scanln(&num1)
	fmt.Println("Enter second Number: ")
	fmt.Scanln(&num2)

	return num1, num2
}

func add(num1 int, num2 int) int {

	return num1 + num2
}

func display(name string, sum int) {

	fmt.Println("Hello, ", name)
	fmt.Println("Summation =", sum)

}

func printGoodBye() {

	fmt.Println("Thank you all")
	fmt.Println("Allah Hafez")

}
func main() {

	printWelcomeMsg()
	name := getUserName()
	num1, num2 := getTwoNumbers()

	sum := add(num1, num2)
	//Display results
	display(name, sum)
	//Print a goodbye message
	printGoodBye()
}

//* SOLID:
/*

S =Single Responsibility Principle


*/
