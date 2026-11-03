package main

import "fmt"

// * function expression or Assign function in Variable
func main() {
	add(2, 3) //don't support hoisting

	add := func(a, b int) {

		c := a + b
		fmt.Println((c))
	}

	sum(2, 3) // support hoisting

}

func sum(a, b int) {
	c := a + b
	fmt.Println((c))
}

func init() {

	fmt.Println("I will be called first")

}

/*




NOTE:

১. ফাংশনের ভেতর Named Function কেন সম্ভব নয়?

গো ল্যাঙ্গুয়েজের ডিজাইন বা সিনট্যাক্স অনুযায়ী, কোনো ফাংশনের স্কোপের (Scope) ভেতরে আরেকটি ফাংশন ব্লক func name() {} এভাবে তৈরি করা যায় না। এটি করলে কম্পাইলার এরর দেবে।
go
func main() {
     এটি ভুল! কম্পাইল এরর আসবে (syntax error: unexpected sum inside function)
    func sum(a, b int) {
        fmt.Println(a + b)
    }
}
Use code with caution.

২. Anonymous Function কীভাবে কাজ করে?

Named function সরাসরি তৈরি করা না গেলেও, আপনি ফাংশনের ভেতরে একটি ফাংশন তৈরি করে সেটিকে একটি ভেরিয়েবলের মধ্যে জমা (Assign) রাখতে পারেন। যেহেতু এর কোনো নির্দিষ্ট নাম থাকে না (ভেরিয়েবলের নামটাই তার পরিচয় হয়), তাই একে Anonymous Function বলে।
go
func main() {
     এটি একদম সঠিক! ফাংশনটিকে 'mysum' ভেরিয়েবলে রাখা হয়েছে
    mysum := func(a, b int) {
        fmt.Println(a + b)
    }

    mysum(5, 5) // আউটপুট: 10
}
Use code with caution.

৩. IIFE (Immediately Invoked Function Expression) কী এবং কীভাবে কাজ করে?

IIFE হলো এমন একটি Anonymous Function যা তৈরি করার সাথে সাথেই নিজে নিজে রান বা কল হয়ে যায়। একে আলাদা করে কোনো ভেরিয়েবলে রাখার বা পরে কল করার প্রয়োজন হয় না।
গো-তে এটি লেখার নিয়ম হলো: ফাংশনের বডির শেষে } এর পরপরই ফার্স্ট ব্র্যাকেট () দিয়ে আর্গুমেন্ট পাস করে দিতে হয়।
go
func main() {
     IIFE উদাহরণ: তৈরি হওয়ার সাথে সাথেই রান হয়ে যাবে
    func(msg string) {
        fmt.Println("হ্যালো,", msg)
    }("বাংলাদেশ") // এখানে ব্র্যাকেট দিয়ে সাথে সাথে কল করা হলো

     আউটপুট আসবে: হ্যালো, বাংলাদেশ
}
Use code with caution.

কেন গো-তে এই নিয়ম রাখা হয়েছে?

১. Closure (ক্লোজার) সুবিধা: ফাংশনের ভেতরে Anonymous function তৈরি করার মূল কারণ হলো ক্লোজার ব্যবহার করা। এর মাধ্যমে ভেতরের ফাংশনটি তার বাইরের ফাংশনের ভেরিয়েবলগুলোকে অ্যাক্সেস এবং মডিফাই করতে পারে।
২. সহজ ও পরিষ্কার কোড: ফাংশনের ভেতরে আবার নতুন Named function তৈরি করতে দিলে কোডের রিডাবিলিটি বা পঠনযোগ্যতা নষ্ট হয় এবং স্কোপিং জটিল হয়ে পড়ে।
আপনার ১ নম্বর প্রশ্নের উত্তর তো পরিষ্কার হলো। আপনার পরের বিষয় বা ২ নম্বর প্রশ্নটি কী? সেটি জানালে আমি একইভাবে সহজ করে বুঝিয়ে দেব।



বৈশিষ্ট্য	Function Declaration

Hoisting (আগে কল করা)	হ্যাঁ, করা যায়।
(ফাইলের যেকোনো জায়গায় ফাংশনটি থাকলেই হলো, উপরে কল করলেও কাজ করে)

Function Expression
না, করা যায় না। (অবশ্যই আগে ভেরিয়েবল তৈরি করতে হবে, তারপর তার নিচে কল করতে হবে) [1, 2]


*/
