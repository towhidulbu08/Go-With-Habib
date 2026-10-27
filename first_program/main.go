package main

import "fmt"

func add(num1 int, num2 int) int {
	return num1 + num2
}

func printSomething()  {
	fmt.Println("Education Must be free")
	
}

func sayHello(name string){

	fmt.Println("Welcome to the golang course",name)

}

func main() {

	printSomething()
	sayHello("Tamim")
}