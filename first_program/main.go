package main

import "fmt"

// হায়ার-অর্ডার ফাংশন (যা প্যারামিটার হিসেবে অন্য ফাংশন নেয়)
// এখানে callback হলো এমন একটি ফাংশন যা int নেয় এবং int রিটার্ন করে
func processNumber(num int, callback func(int) int) int {
	return callback(num)
}

func main() {
	// একটি অ্যানোনিমাস বা নামহীন ফাংশন তৈরি করে পাস করা হলো (যা স্কয়ার করে)
	square := func(n int) int {
		return n * n
	}

	result := processNumber(5, square)
	fmt.Println(result) // আউটপুট: 25
}

/*

//*Parameter Vs Argument
//*First Order Function
//*Higher Order Function and First Class Function
//*Callback Function


*/
