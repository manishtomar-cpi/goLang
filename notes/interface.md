# Interfaces in Go

## What is an Interface?

An **interface** in Go is a **type** that defines a set of method signatures.  
If a type implements all the methods in an interface, it **automatically satisfies** that interface.  
This allows **polymorphism** — writing code that works with many types, as long as they implement the same behavior.

---

## Basic Example

```go
package main

import "fmt"

// Define an interface
type Speaker interface {
    Speak()
}

// Define structs
type Dog struct{}
func (d Dog) Speak() {
    fmt.Println("Woof!")
}

type Cat struct{}
func (c Cat) Speak() {
    fmt.Println("Meow!")
}

// Function that accepts any type implementing Speaker
func makeItSpeak(s Speaker) {
    s.Speak()
}

func main() {
    dog := Dog{}
    cat := Cat{}

    makeItSpeak(dog)
    makeItSpeak(cat)
}
```

**Explanation:**
- `Speaker` defines one method: `Speak()`.
- Both `Dog` and `Cat` have a `Speak()` method.
- Therefore, they **satisfy the Speaker interface** automatically.
- The function `makeItSpeak()` can accept any type that implements `Speak()`.

---

## Key Point: Implicit Implementation

In Go, a type **does not need to explicitly declare** that it implements an interface.  
If it has all the required methods, it satisfies the interface automatically.

This is called **structural typing**.

---

## Empty Interface (`interface{}`)

An **empty interface** has **no methods**, so **every type** satisfies it.  
It can hold a value of **any type**.

```go
func describe(i interface{}) {
    fmt.Printf("Type: %T, Value: %v\n", i, i)
}

func main() {
    describe(42)
    describe("hello")
    describe(true)
}
```

Output:
```
Type: int, Value: 42
Type: string, Value: hello
Type: bool, Value: true
```

---

## Type Assertion

Type assertion extracts the actual type stored in an interface variable.

```go
var s Speaker = Dog{}
dog := s.(Dog) // Type assertion
dog.Speak()
```

If the type does not match, Go will panic at runtime.  
You can prevent this using the **comma-ok** form:

```go
dog, ok := s.(Dog)
if ok {
    dog.Speak()
}
```

---

## Type Switch

Use a **type switch** to handle multiple types stored in an interface.

```go
switch v := s.(type) {
case Dog:
    fmt.Println("This is a Dog")
case Cat:
    fmt.Println("This is a Cat")
default:
    fmt.Println("Unknown type")
}
```

---

## Why Use Interfaces?

- To make code **flexible** and **reusable**
- To **decouple** behavior from implementation
- To **mock types easily** for testing
- To design **clean, abstract APIs**

Example:

```go
type Writer interface {
    Write(data []byte) (int, error)
}

func saveFile(w Writer, data []byte) {
    w.Write(data)
}
```

Any type with a `Write()` method (like `os.File`) can be used with `saveFile`.

---

## Summary Table

| Concept | Description |
|----------|--------------|
| Definition | Collection of method signatures |
| Implementation | Implicit (no need to declare `implements`) |
| Supports | Polymorphism |
| Empty Interface | Accepts any type |
| Extract type | Type assertion or type switch |
| Common use | Abstraction, testing, flexibility |

---

## Summary

Interfaces in Go define behavior, not data.  
Any type that implements all the methods of an interface automatically satisfies it.  
This enables writing flexible, reusable, and decoupled code without traditional class-based inheritance.
