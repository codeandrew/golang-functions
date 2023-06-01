# Go Lang

## Intro

```go
// Package Declaration
package main

// Import
import "fmt"
// fmt shorthand for format 

// Println means Print Line
func main(){
  fmt.Println("Hello World")
}
// go run helloworld.go
```

### Boolean
```go
func main() {
  fmt.Println(true && true)
  fmt.Println(true && false)
  fmt.Println(true || true)
  fmt.Println(true || false)
  fmt.Println(!true)
}
// $ go run main.go
// true
// false
// true
// true
// false

```
### Variables


In Go, variables are used to store values. This guide explains different ways to declare variables in Go.

1. Basic Variable Declaration

You can declare a variable with `var`, then the variable name, and its type:

```go
var age int
```

Here, `age` is a variable of type `int`.

2. Declare and Initialize

You can also declare a variable and assign it a value at the same time:

```go
var age int = 25
```

3. Type Inference

If you assign a value to the variable when you declare it, Go can figure out the type for you:

```go
var age = 25 // Go infers `age` is an int
```

4. Short Declaration

For a shorter syntax, use `:=` to declare and initialize a variable:

```go
age := 25 // Go infers `age` is an int
```
Please note, the short declaration can only be used inside functions.

5. Multiple Variables

You can declare and initialize multiple variables in one line:

```go
var length, width int = 20, 30
```

Or with short declaration:

```go
length, width := 20, 30
```

You can also use the `var()` block to declare multiple variables of different types:

```go
var (
    name = "John"
    age = 30
)
```



## References
- https://www.golang-book.com/books/intro/2
