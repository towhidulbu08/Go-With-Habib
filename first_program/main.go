package main

import "fmt"

var a = 4

func main() {
	fmt.Println(a) //output:a=20
}

func init() {

	fmt.Println(a) //output:a=4
	a = 20
}

//*Function Types: Init function

//NOTE: init function is executed before main function
