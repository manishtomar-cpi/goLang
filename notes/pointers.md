# Go Pointers - Simple Notes

## What is a Pointer?
A pointer is a variable that stores the **memory address** of another variable.
Instead of holding the actual value, it points to where the value is stored in memory.

### Example
```go
package main
import "fmt"

func main() {
    x := 10
    p := &x // p stores the address of x

    fmt.Println("x =", x)
    fmt.Println("p =", p)   // memory address
    fmt.Println("*p =", *p) // value at that address
}
```
- `&x` gives the address of `x`.
- `*p` gives the value stored at that address.

---

## Why Use Pointers?
In Go, function arguments are passed **by value** (a copy is made).  
If you want to modify the original variable, you need to pass a **pointer**.

### Example Without Pointer
```go
func changeValue(val int) {
    val = 100
}

func main() {
    x := 10
    changeValue(x)
    fmt.Println(x) // still 10, not changed
}
```

### Example With Pointer
```go
func changeValue(val *int) {
    *val = 100
}

func main() {
    x := 10
    changeValue(&x)
    fmt.Println(x) // now 100
}
```

---

## When to Use Pointers

### 1. Update Struct Fields in Functions
Useful when you want to modify data inside a struct directly.
```go
type Device struct {
    Name  string
    Power bool
}

func turnOn(d *Device) {
    d.Power = true
}

func main() {
    dev := Device{Name: "Sensor1", Power: false}
    turnOn(&dev)
    fmt.Println(dev.Power) // true
}
```

Use Case: Updating state in APIs, IoT devices, or configurations.

---

### 2. Avoid Copying Large Data
Passing large data structures by pointer saves memory.
```go
type Image struct {
    Pixels [1000000]int
}

func process(img *Image) {
    // process without copying large data
}
```

Use Case: Image processing or large datasets.

---

### 3. Linked Data Structures
Pointers connect elements like in linked lists or trees.
```go
type Node struct {
    value int
    next  *Node
}

func main() {
    n1 := &Node{value: 10}
    n2 := &Node{value: 20}
    n1.next = n2
    fmt.Println(n1.next.value) // 20
}
```

Use Case: Implementing data structures such as queues, lists, and graphs.

---

### 4. Interacting with Hardware or External Libraries
Pointers are used when working with memory addresses or embedded devices.
```go
func readSensor(data *int) {
    *data = 42 // pretend reading from sensor
}
```

Use Case: Embedded systems and IoT communication.

---

### 5. Methods with Pointer Receivers
Allows methods to modify struct fields directly.
```go
type Account struct {
    balance int
}

func (a *Account) Deposit(amount int) {
    a.balance += amount
}

func main() {
    acc := Account{balance: 100}
    acc.Deposit(50)
    fmt.Println(acc.balance) // 150
}
```

Use Case: Updating object state (e.g., in banking or inventory systems).

---

## When Not to Use Pointers
Avoid pointers when:
- The data is small and does not need to be modified.
- The function only needs to read the data.
- Pointers make the code harder to read without any real benefit.

---

## Summary

| Concept | Meaning | Example |
|----------|----------|----------|
| `&x` | Address of variable `x` | `p := &x` |
| `*p` | Value at the pointer | `fmt.Println(*p)` |
| Pass by value | Copies the variable | `changeValue(x)` |
| Pass by reference | Modifies the original value | `changeValue(&x)` |
| Pointer receiver | Allows struct modification | `func (a *Account) Deposit()` |

---

## Simple Understanding
Pointers help you share and modify data directly in memory.  
They are useful for improving performance, updating data inside functions, and managing large or complex structures efficiently.
