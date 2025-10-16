# Difference Between `var` and `:=` in Go

## 1. Using `var`
- Used for declaring variables (with or without initial values).
- Can be used outside or inside functions.
- You can specify the type explicitly.
- If you don’t give a value, Go assigns the zero value for that type.

### Example 1: Declare with type
```go
var name string = "Manish"
var age int = 26
```

### Example 2: Declare without value (uses zero value)
```go
var number int
fmt.Println(number) // Output: 0
```

### Example 3: Declare without specifying type (type inferred)
```go
var city = "Delhi"
fmt.Println(city) // Output: Delhi
```

---

## 2. Using `:=`
- Called the short variable declaration.
- Used only inside functions.
- Go automatically infers the type from the assigned value.
- You must assign a value at the time of declaration.

### Example:
```go
name := "Manish"
age := 26
fmt.Println(name, age)
```

---

## 3. Key Differences

| Feature | `var` | `:=` |
|----------|--------|-------|
| Can be used outside functions | Yes | No |
| Requires initial value | No | Yes |
| Type specification | Optional | Not allowed (auto-inferred) |
| Re-declaration | Possible in new scope | Not allowed for same variable in same scope |
| Common usage | Package-level or global vars | Inside functions for quick declarations |

---

## 4. Example Showing Both
```go
package main

import "fmt"

var globalVar = "I am global" // using var (allowed outside functions)

func main() {
    localVar := "I am local"   // using :=
    fmt.Println(globalVar)
    fmt.Println(localVar)
}
```

Output:
```
I am global
I am local
```

---

## Summary
- Use `var` when declaring global or zero-value variables.
- Use `:=` inside functions for quick variable creation.
