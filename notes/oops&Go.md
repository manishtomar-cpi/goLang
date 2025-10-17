# Is Go Object-Oriented?

Go (Golang) supports several object-oriented programming (OOP) concepts but does not follow traditional OOP as seen in languages like Java or C++.

## OOP Features Go Supports

### 1. Encapsulation
Go provides encapsulation through exported and unexported identifiers.
- Identifiers starting with a capital letter are **exported** (public).
- Identifiers starting with a lowercase letter are **unexported** (private).

Example:
```go
type Person struct {
    Name string
    age  int // unexported (private)
}
```

### 2. Composition Instead of Inheritance
Go does not support classical inheritance. Instead, it uses **struct embedding** for code reuse and composition.

Example:
```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println("Some sound")
}

type Dog struct {
    Animal // embedded struct
}

func (d Dog) Speak() {
    fmt.Println("Woof!")
}
```

Here, `Dog` embeds `Animal` and can access its methods but is not a subclass.

### 3. Polymorphism Through Interfaces
Go provides polymorphism using **interfaces**, which are implemented implicitly (no need for an explicit `implements` keyword).

Example:
```go
type Speaker interface {
    Speak()
}

func MakeItSpeak(s Speaker) {
    s.Speak()
}
```

Any type that defines a `Speak()` method automatically satisfies the `Speaker` interface.

## Features Not Supported in Go

- No classes
- No inheritance hierarchy
- No constructors or destructors
- No method overloading
- No operator overloading

## Summary Table

| OOP Concept     | Supported in Go | How                                         |
|------------------|----------------|---------------------------------------------|
| Encapsulation    | Yes             | Exported/unexported identifiers             |
| Inheritance      | No              | Use composition (struct embedding)          |
| Polymorphism     | Yes             | Interfaces                                 |
| Abstraction      | Yes             | Interfaces and packages                     |
| Classes          | No              | Use structs with methods                    |

## Summary

Go is object-oriented in a minimal and composition-based way. It supports organization, reuse, and polymorphism without the complexity of classical OOP inheritance chains.
