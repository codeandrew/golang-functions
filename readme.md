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
### Control Structures

Control structures in Go allow you to control the flow of your program. Here are the basics:

**1. If-Else**

You can use `if` to execute a block of code only if a certain condition is true. `else` can be used to specify a block of code to be executed if the same condition is false.

```go
if x > 0 {
    fmt.Println("x is positive")
} else {
    fmt.Println("x is not positive")
}
```

**2. For**

`for` is used for looping. The basic `for` loop has three components separated by semicolons: the init statement, the condition expression, and the post statement.

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**3. Switch**

`switch` is a multi-way branch statement. It can be more efficient and easier to read than long `if-else` chains.

```go
switch day {
case "Monday":
    fmt.Println("Start of the week")
case "Friday":
    fmt.Println("End of the week")
default:
    fmt.Println("Midweek")
}
```

**Best Practices**

1. Keep `if` conditions simple; complex logic should be moved to functions.
2. Limit the use of `else`; try to return early from functions instead.
3. Prefer `for` over `while` for clarity.
4. Use `switch` instead of multi-way `if-else`.
5. Always have a `default` case in `switch` to handle unexpected cases.

Here's a function implementing these practices:

```go
func getDayType(day string) string {
    if day == "" {
        return "Unknown"
    }

    switch day {
    case "Monday":
        return "Start of the week"
    case "Friday":
        return "End of the week"
    default:
        return "Midweek"
    }
}
```

In this function, input is checked with `if`, and the flow is controlled with `switch`. The function returns as soon as it knows the result, which keeps the code clear and easy to read.


## References
- https://www.golang-book.com/books/intro/2
