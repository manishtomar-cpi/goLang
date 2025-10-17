# Go Functions - Simple Notes

## What is a Function?
A function in Go is a block of code that performs a specific task.  
Functions make code reusable and easier to organize.

### Syntax
```go
func functionName(parameters) returnType {
    // code to execute
}
```

### Example
```go
func greet() {
    fmt.Println("Hello, Go!")
}
```

---

## Function with Parameters
Functions can take input parameters.

```go
func add(a int, b int) int {
    return a + b
}
```
If all parameters are of the same type, you can write:
```go
func add(a, b int) int {
    return a + b
}
```

---

## Function Returning Multiple Values
Go functions can return more than one value.

```go
func operations(a, b int) (int, int, int) {
    return a + b, a - b, a * b
}

sum, diff, prod := operations(5, 3)
fmt.Println(sum, diff, prod) // 8 2 15
```

If you do not need all return values, use `_`:
```go
sum, _, prod := operations(5, 3)
```

---

## Variadic Functions
A variadic function can take any number of arguments.

```go
func sumAll(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

fmt.Println(sumAll(1, 2, 3, 4)) // 10
```

### Use Case
Variadic functions are useful when the number of inputs is not fixed.  
For example, when adding multiple numbers, calculating averages, logging multiple messages, or formatting strings.  
Instead of defining many parameters, we can just use `...` to handle all arguments flexibly.

---

## Function as a Parameter (First-Class Function)
In Go, functions can be passed to other functions as arguments.

```go
func process(fn func(int) int) {
    fmt.Println(fn(5))
}

func main() {
    square := func(x int) int { return x * x }
    process(square) // prints 25
}
```

### Use Case
First-class functions are useful for:
- Passing custom logic to another function (for example, filters or sorting conditions)
- Creating middlewares in web servers (logging, authentication)
- Implementing dynamic behavior such as different payment methods, retry logic, or event handling

---

## Returning a Function
Functions can return another function.

```go
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

double := multiplier(2)
fmt.Println(double(5)) // 10
```

---

## Anonymous Functions
Functions without a name are called anonymous functions.

```go
func() {
    fmt.Println("Hello from anonymous function")
}()
```
The `()` at the end calls it immediately.

---

## Summary Table

| Concept | Description | Example |
|----------|--------------|----------|
| Regular Function | Named block of code | `func greet() {}` |
| Parameters | Inputs to a function | `func add(a, b int)` |
| Multiple Returns | Return more than one value | `(int, int)` |
| Variadic Function | Takes many arguments | `func sumAll(nums ...int)` |
| Anonymous Function | Function without a name | `func() {}` |
| Function as Parameter | Pass function to another function | `process(square)` |
| Returning Function | Return logic dynamically | `multiplier(2)` |

---

## Complete Example
```go
package main

import "fmt"

func sumAll(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

func process(fn func(int) int) {
    fmt.Println("Result:", fn(5))
}

func main() {
    fmt.Println("Sum:", sumAll(1, 2, 3, 4))

    square := func(x int) int { return x * x }
    process(square)
}
```
