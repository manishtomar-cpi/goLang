# Struct Embedding vs Inheritance in Go

## Is Struct Embedding Pure Inheritance?

No, **struct embedding in Go is not pure inheritance**.  
It may look similar because it allows code reuse, but it is actually based on **composition**, not inheritance.

---

## 1. What Is Struct Embedding?

**Struct embedding** allows one struct to include another struct as a field.  
This gives the outer struct access to the fields and methods of the embedded struct directly.

### Example

```go
package main

import "fmt"

type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
    Animal // embedded struct
    Breed  string
}

func main() {
    d := Dog{Animal: Animal{Name: "Buddy"}, Breed: "Labrador"}
    d.Speak() // Accessible directly through embedding
}
```

Here:
- `Dog` contains an `Animal` (it **has an** Animal).
- There is **no parent-child relationship** — this is **composition**, not inheritance.

---

## 2. What Is Inheritance?

In traditional object-oriented programming (OOP), **inheritance** means one class derives from another and automatically gets its properties and methods.

### Example (Java)

```java
class Animal {
    void speak() {
        System.out.println("Animal speaks");
    }
}

class Dog extends Animal {
    void bark() {
        System.out.println("Dog barks");
    }
}
```

Here:
- `Dog` **is an** `Animal`.  कुत्ता एक जानवर है
- This forms a **parent-child hierarchy**.
- `Dog` inherits all the behavior of `Animal`.

---

## 3. Key Differences Between Inheritance and Embedding

| Feature | Classical Inheritance (Java/C++) | Struct Embedding (Go) |
|----------|----------------------------------|------------------------|
| Relationship | “Is-a” (Dog **is an** Animal) | “Has-a” (Dog **has an** Animal) कुत्ते के पास एक जानवर है |
| Mechanism | Parent-child hierarchy | Composition (embedding) |
| Code Reuse | Through inheritance hierarchy | Through delegation (method promotion) |
| Overriding | Subclass can override parent methods | Embedded type’s methods can be shadowed |
| Polymorphism | Based on type hierarchy | Based on interfaces |
| Multiple inheritance | Usually not allowed | Supported via multiple embedding |

---

## 4. Method Promotion vs Inheritance

When you embed a struct, Go **promotes** the embedded struct’s methods to the outer struct automatically.

### Example

```go
type Animal struct {
    Name string
}

func (a Animal) Speak() {
    fmt.Println(a.Name, "makes a sound")
}

type Dog struct {
    Animal
}

func (d Dog) Speak() {
    fmt.Println(d.Name, "barks")
}

func main() {
    d := Dog{Animal: Animal{Name: "Rocky"}}
    d.Speak()         // Dog’s Speak() shadows Animal’s Speak()
    d.Animal.Speak()  // Still accessible
}
```

This is **method shadowing**, not classical method overriding.

---

## 5. Why Go Uses Composition, Not Inheritance

Go avoids classical inheritance because it:
- Simplifies the type system (no complex class trees)
- Reduces coupling between types
- Encourages **composition over inheritance**

Go’s philosophy:
> “Don’t build deep type hierarchies; build simple, composable systems.”

---

## 6. Summary

| Concept | Go’s Struct Embedding | OOP Inheritance |
|----------|------------------------|----------------|
| Relationship | Has-a | Is-a |
| Reuse mechanism | Composition | Hierarchy |
| Coupling | Loose | Tight |
| Flexibility | High | Limited |
| Polymorphism | Via interfaces | Via class hierarchy |
| Supported in Go? |  Yes |  No (directly) |

---

## In One Line

> **Struct embedding in Go is composition, not pure inheritance.**  
> It enables code reuse and method promotion without parent-child hierarchies or class inheritance.
