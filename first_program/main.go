package main

import "fmt"

// হায়ার-অর্ডার ফাংশন (যা নিজে একটি ফাংশন রিটার্ন করে)
func multiplier(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}

func main() {
	double := multiplier(2)
	fmt.Println(double(10)) // আউটপুট: 20

	// multiplier(3) আরেকটি ফাংশন রিটার্ন করবে যা সংখ্যাকে ৩ দিয়ে গুণ করে
	triple := multiplier(3)
	fmt.Println(triple(10)) // আউটপুট: 30
}

/*
First Class Citizen: A first-class citizen (or first-class element) is any entity in a programming language that supports all standard operations available to other entities. This is a broad language design concept, not just a function concept.
If something is a first-class citizen, it can be:
• Stored in a variable or data structure.
• Passed as an argument to a function.
• Returned from a function.
• Created dynamically at runtime

First-Class Function:

A language has first-class functions when functions are treated like first-class citizens. This means the language treats functions just like regular data (like integers or strings).

Higher-Order Function: The function who takes a function as an argument or returns a function.


First-Order Function:

A first-order function is a standard, regular function that only accepts data as inputs and only returns data as an output.
• It cannot accept another function as an argument.
• It cannot return a function.
• It simply processes raw values (like numbers, strings, or objects).

Callback Function: A callback function is a function that is passed as an argument to another function and is executed after some operation has been completed. It allows for asynchronous programming and event handling.

*/
