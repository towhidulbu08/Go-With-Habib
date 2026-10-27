# Golang Learning Guide

A structured, beginner-friendly guide to learning Go (Golang) step by step, with code examples and practice tasks.

---

## Table of Contents

1. [Introduction](#1-introduction)
2. [Setup](#2-setup)
3. [Hello World](#3-hello-world)
4. [Variables and Data Types](#4-variables-and-data-types)
5. [Control Flow](#5-control-flow)
6. [Functions](#6-functions)
7. [Arrays, Slices, and Maps](#7-arrays-slices-and-maps)
8. [Structs and Methods](#8-structs-and-methods)
9. [Pointers](#9-pointers)
10. [Interfaces](#10-interfaces)
11. [Error Handling](#11-error-handling)
12. [Goroutines and Channels](#12-goroutines-and-channels)
13. [Packages and Modules](#13-packages-and-modules)
14. [Testing](#14-testing)
15. [Project Ideas](#15-project-ideas)
16. [Learning Roadmap](#16-learning-roadmap)
17. [Resources](#17-resources)

---

## 1. Introduction

**Go** is an open-source language created at Google. It is known for:

- Simple, clean syntax
- Fast compilation and execution
- Built-in concurrency (goroutines)
- Strong standard library
- Single binary output, easy to deploy

**Common uses:** backend APIs, CLI tools, cloud and DevOps tools (Docker, Kubernetes), microservices.

---

## 2. Setup

1. Download Go from [go.dev/dl](https://go.dev/dl/)
2. Verify the installation:

```bash
go version
```

3. Create your first project:

```bash
mkdir go-learning
cd go-learning
go mod init go-learning
```

### Useful commands

| Command         | Purpose                           |
| --------------- | --------------------------------- |
| `go run .`      | Compile and run                   |
| `go build`      | Build a binary                    |
| `go fmt ./...`  | Format all code                   |
| `go vet ./...`  | Check for suspicious code         |
| `go test ./...` | Run all tests                     |
| `go mod tidy`   | Clean up dependencies             |

---

## 3. Hello World

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

**Key points**

- Every file starts with `package`
- `package main` + `func main()` is the entry point
- `import` brings in other packages

---

## 4. Variables and Data Types

### Declaring variables

```go
var name string = "Rahim"
var age int = 22

city := "Dhaka"        // short declaration (inside functions only)
const pi = 3.14159     // constant
```

### Basic types

| Type                         | Example              |
| ---------------------------- | -------------------- |
| `int`, `int8` ... `int64`    | `42`                 |
| `uint`, `uint8` ... `uint64` | `10`                 |
| `float32`, `float64`         | `3.14`               |
| `string`                     | `"hello"`            |
| `bool`                       | `true`, `false`      |
| `byte` (alias of `uint8`)    | `'a'`                |
| `rune` (alias of `int32`)    | `'অ'`                |

### Zero values

Variables declared without a value get a default:

```go
var a int      // 0
var b string   // ""
var c bool     // false
```

### Type conversion

```go
x := 10
y := float64(x)
```

---

## 5. Control Flow

### if / else

```go
if age >= 18 {
    fmt.Println("Adult")
} else if age >= 13 {
    fmt.Println("Teen")
} else {
    fmt.Println("Child")
}
```

### switch

```go
switch day {
case "Sat", "Sun":
    fmt.Println("Weekend")
default:
    fmt.Println("Weekday")
}
```

### for (the only loop in Go)

```go
// classic
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// while-style
for count < 10 {
    count++
}

// infinite
for {
    break
}

// range
for index, value := range []string{"a", "b", "c"} {
    fmt.Println(index, value)
}
```

**Practice:** Write FizzBuzz for numbers 1 to 100.

---

## 6. Functions

```go
func add(a int, b int) int {
    return a + b
}
```

### Multiple return values

```go
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("cannot divide by zero")
    }
    return a / b, nil
}
```

### Variadic functions

```go
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}
```

### Anonymous functions and closures

```go
counter := func() func() int {
    c := 0
    return func() int {
        c++
        return c
    }
}()

fmt.Println(counter()) // 1
fmt.Println(counter()) // 2
```

### defer

`defer` runs a call when the surrounding function returns.

```go
func readFile() {
    f, _ := os.Open("data.txt")
    defer f.Close()
    // work with f
}
```

---

## 7. Arrays, Slices, and Maps

### Arrays (fixed size)

```go
var arr [3]int = [3]int{1, 2, 3}
```

### Slices (dynamic, most commonly used)

```go
nums := []int{1, 2, 3}
nums = append(nums, 4, 5)

fmt.Println(len(nums), cap(nums))
fmt.Println(nums[1:3]) // [2 3]

copyNums := make([]int, len(nums))
copy(copyNums, nums)
```

### Maps (key-value)

```go
ages := map[string]int{
    "Rahim": 22,
    "Karim": 25,
}

ages["Salma"] = 20
delete(ages, "Karim")

if age, ok := ages["Rahim"]; ok {
    fmt.Println("Found:", age)
}
```

> **Remember:** Writing to a `nil` map causes a panic. Always create maps with `make` or a literal.

**Practice:** Count word frequency in a sentence using a map.

---

## 8. Structs and Methods

```go
type Person struct {
    Name string
    Age  int
}

// Method with value receiver
func (p Person) Greet() string {
    return "Hi, I am " + p.Name
}

// Method with pointer receiver (can modify the struct)
func (p *Person) Birthday() {
    p.Age++
}
```

```go
p := Person{Name: "Rahim", Age: 22}
fmt.Println(p.Greet())
p.Birthday()
```

### Embedding (composition)

```go
type Employee struct {
    Person
    Role string
}
```

---

## 9. Pointers

A pointer stores the **memory address** of a value.

```go
x := 10
p := &x        // address of x
fmt.Println(*p) // 10 (dereference)
*p = 20
fmt.Println(x)  // 20
```

**Why use pointers?**

- Modify the original value inside a function
- Avoid copying large structs

```go
func increment(n *int) {
    *n++
}
```

---

## 10. Interfaces

An interface defines a set of methods. Any type that has those methods satisfies it automatically (no `implements` keyword).

```go
type Shape interface {
    Area() float64
}

type Rectangle struct{ W, H float64 }
type Circle struct{ R float64 }

func (r Rectangle) Area() float64 { return r.W * r.H }
func (c Circle) Area() float64    { return math.Pi * c.R * c.R }

func printArea(s Shape) {
    fmt.Println(s.Area())
}
```

### Empty interface and type assertion

```go
var v any = "hello"

s, ok := v.(string)
if ok {
    fmt.Println(s)
}
```

---

## 11. Error Handling

Go uses explicit error values instead of exceptions.

```go
result, err := divide(10, 0)
if err != nil {
    fmt.Println("Error:", err)
    return
}
fmt.Println(result)
```

### Custom errors

```go
type ValidationError struct {
    Field string
}

func (e ValidationError) Error() string {
    return "invalid field: " + e.Field
}
```

### Wrapping errors

```go
if err != nil {
    return fmt.Errorf("load config: %w", err)
}
```

```go
if errors.Is(err, os.ErrNotExist) {
    // handle missing file
}
```

### panic and recover

Use `panic` only for truly unrecoverable situations.

```go
defer func() {
    if r := recover(); r != nil {
        fmt.Println("Recovered:", r)
    }
}()
```

---

## 12. Goroutines and Channels

### Goroutines

```go
go func() {
    fmt.Println("running concurrently")
}()
```

### WaitGroup

```go
var wg sync.WaitGroup

for i := 1; i <= 3; i++ {
    wg.Add(1)
    go func(id int) {
        defer wg.Done()
        fmt.Println("worker", id)
    }(i)
}

wg.Wait()
```

### Channels

```go
ch := make(chan string)

go func() {
    ch <- "ping"
}()

msg := <-ch
fmt.Println(msg)
```

### Buffered channels and range

```go
jobs := make(chan int, 3)
jobs <- 1
jobs <- 2
jobs <- 3
close(jobs)

for j := range jobs {
    fmt.Println(j)
}
```

### select

```go
select {
case msg := <-ch1:
    fmt.Println(msg)
case <-time.After(time.Second):
    fmt.Println("timeout")
}
```

### Mutex (protect shared data)

```go
var (
    mu      sync.Mutex
    counter int
)

mu.Lock()
counter++
mu.Unlock()
```

> Check for data races with: `go run -race main.go`

---

## 13. Packages and Modules

### Project layout

```
myapp/
├── go.mod
├── main.go
└── mathutil/
    └── mathutil.go
```

### Creating a package

```go
// mathutil/mathutil.go
package mathutil

func Add(a, b int) int { // capital letter = exported
    return a + b
}
```

### Using it

```go
import "myapp/mathutil"

fmt.Println(mathutil.Add(2, 3))
```

### Adding external packages

```bash
go get github.com/google/uuid
go mod tidy
```

**Rule:** Names starting with a **capital letter** are exported (public). Lowercase names are private to the package.

---

## 14. Testing

Test files end with `_test.go`.

```go
// mathutil/mathutil_test.go
package mathutil

import "testing"

func TestAdd(t *testing.T) {
    got := Add(2, 3)
    want := 5
    if got != want {
        t.Errorf("Add(2, 3) = %d; want %d", got, want)
    }
}
```

### Table-driven tests

```go
func TestAddTable(t *testing.T) {
    tests := []struct {
        name string
        a, b int
        want int
    }{
        {"positive", 2, 3, 5},
        {"negative", -1, -2, -3},
        {"zero", 0, 0, 0},
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            if got := Add(tt.a, tt.b); got != tt.want {
                t.Errorf("got %d, want %d", got, tt.want)
            }
        })
    }
}
```

Run with:

```bash
go test ./... -v
```

---

## 15. Project Ideas

| Level        | Project                                   | Skills practiced                  |
| ------------ | ----------------------------------------- | --------------------------------- |
| Beginner     | Number guessing game                      | loops, input, conditions          |
| Beginner     | Temperature converter                     | functions, types                  |
| Beginner     | To-do list (CLI)                          | slices, structs, file I/O         |
| Intermediate | Contact book with JSON storage            | maps, JSON, error handling        |
| Intermediate | Simple REST API (`net/http`)              | HTTP, routing, JSON               |
| Intermediate | Concurrent web scraper                    | goroutines, channels, WaitGroup   |
| Advanced     | URL shortener with database               | SQL, REST, testing                |
| Advanced     | Chat server with WebSockets               | concurrency, networking           |

---

## 16. Learning Roadmap

**Week 1: Basics**
- [ ] Install Go, run Hello World
- [ ] Variables, types, constants
- [ ] if, switch, for

**Week 2: Core**
- [ ] Functions, multiple returns, defer
- [ ] Arrays, slices, maps
- [ ] Structs and methods

**Week 3: Intermediate**
- [ ] Pointers
- [ ] Interfaces
- [ ] Error handling

**Week 4: Concurrency and Tools**
- [ ] Goroutines, channels, select
- [ ] Packages and modules
- [ ] Testing

**Week 5+: Build**
- [ ] Build a CLI app
- [ ] Build a REST API
- [ ] Read other people's Go code

---

## 17. Resources

- [A Tour of Go](https://go.dev/tour/) (interactive, start here)
- [Go by Example](https://gobyexample.com/)
- [Effective Go](https://go.dev/doc/effective_go)
- [Go Standard Library Docs](https://pkg.go.dev/std)
- [Go Playground](https://go.dev/play/) (run code in the browser)
- Book: *The Go Programming Language* by Donovan and Kernighan

---

## Quick Tips

1. Run `go fmt` often. Go has one standard code style.
2. Handle every error. Don't ignore `err`.
3. Keep functions small and packages focused.
4. Prefer composition over inheritance.
5. Write code every day, even for 30 minutes.

**Happy learning, and enjoy Go!**
