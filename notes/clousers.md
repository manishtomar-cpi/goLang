# Go Closures - Simple Notes

## What is a Closure?
A **closure** in Go is a function that remembers variables from the scope in which it was created, even after that scope has finished executing.

In simple terms, when an **inner function** uses a variable from an **outer function**, that variable stays alive as long as the inner function exists.

---

## Example
```go
package main

import "fmt"

func counter() func() int {
    var count int = 0
    return func() int {
        count += 1
        return count
    }
}

func main() {
    increment := counter()
    fmt.Println(increment()) // 1
    fmt.Println(increment()) // 2
}
```

### Explanation Step-by-Step
1. When we call `counter()`, it creates a local variable `count = 0`.
2. The `counter()` function returns another function (the inner one).
3. This inner function still has access to `count` from `counter()`.
4. Even though `counter()` has finished running, `count` is not destroyed.
5. Every time we call `increment()`, it updates the same `count` value.

So the variable `count` stays **alive inside the closure**.

---

## Key Points
- Closures allow functions to **remember state** between calls.
- The variables from the outer function remain accessible even after that function ends.
- Each closure gets its **own copy** of the outer variables.

---

## Another Example
```go
func newIDGenerator() func() int {
    id := 100
    return func() int {
        id++
        return id
    }
}

func main() {
    generateID := newIDGenerator()
    fmt.Println(generateID()) // 101
    fmt.Println(generateID()) // 102

    newGenerator := newIDGenerator()
    fmt.Println(newGenerator()) // 101 (separate copy)
}
```

Here, each call to `newIDGenerator()` creates a **new closure** with its own `id` variable.

---

## Real-World Use Cases of Closures
- **Counters:** To keep track of state between function calls.
- **ID Generators:** To create unique IDs without global variables.
- **Event Handlers:** When you need a function to remember some information.
- **Data Hiding:** To keep data private within a function scope.

---

## Summary
| Concept | Explanation |
|----------|--------------|
| Closure | A function that remembers values from its outer scope |
| Outer Function | The function that defines the variable (e.g., `counter`) |
| Inner Function | The function that uses that variable |
| State Preservation | Variables stay alive as long as the inner function exists |

---

## Simple Understanding
Think of a closure as a **function + its memory**.  
When you call the inner function, it can still access and modify the variables it captured earlier.
