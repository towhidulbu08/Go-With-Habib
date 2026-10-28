package main

import "fmt"

// import "fmt"

var (
	a = 10
	b = 40
)

func add(x int, y int) {

	z := x + y
	fmt.Println(z)
}

func main() {

	p := 30
	q := 40

	add(p, q)
	add(a, b)

	add(b, p)
}
