# Escape Analysis in Go

## What is Escape Analysis?

**Escape analysis** in Go is the process the compiler uses to determine **whether a variable should be stored on the stack or the heap**.

It is part of Go’s **memory management** system. The goal is to allocate variables efficiently and automatically manage their lifetime without manual memory handling.

---

## Stack vs Heap

| Memory Type | Description | Lifetime | Managed By |
|--------------|--------------|-----------|-------------|
| **Stack** | Fast, limited memory area for local variables | Function scope (short-lived) | Automatically freed when function returns |
| **Heap** | Larger, slower memory for dynamic allocations | Exists until garbage collector frees it | Managed by garbage collector |

Normally, local variables are stored on the **stack**.  
But if the compiler detects that a variable’s address **"escapes"** the function (i.e., it’s used outside of it), it gets moved to the **heap**.

---

## Example: Stack Allocation

```go
func createValue() int {
    x := 10
    return x
}
```

Here, `x` is stored on the **stack** because its value is returned directly (not its address).  
Once the function ends, the stack frame is destroyed.

---

## Example: Heap Allocation (Escaping Variable)

```go
func createPointer() *int {
    x := 10
    return &x
}
```

In this case:
- `x` is a local variable.
- You return its **address (`&x`)**.
- The compiler sees that `x` must live beyond the function call.

So, `x` **escapes to the heap**.

---

## Example Output

You can verify this behavior using the **Go compiler flag** `-gcflags`:

```bash
go run -gcflags="-m" main.go
```

Example output:
```
./main.go:4:6: moved to heap: x
```
This means variable `x` was allocated on the heap.

---

## Real-World Example

```go
type student struct {
    name string
}

func newStudent(name string) *student {
    s := student{name: name}
    return &s // s escapes to the heap
}
```

Even though `s` looks local, Go moves it to the **heap** because its pointer is returned.

---

## Why Escape Analysis Matters

1. **Performance**
   - Stack allocations are faster and automatically cleaned up.
   - Heap allocations involve the garbage collector and are slower.

2. **Optimization**
   - Writing code that minimizes unnecessary heap allocations can improve performance.

3. **Safety**
   - Prevents returning pointers to destroyed stack memory.

---

## When a Variable Escapes to the Heap

A variable "escapes" when:
- Its **address** is returned or passed outside the function.
- It’s **captured by a closure** (used inside an anonymous function).
- It’s **stored in an interface** or passed to a function expecting an interface type.
- It’s **stored in a heap-allocated object**.

Example (closure case):
```go
func counter() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}
```
Here, `x` escapes because it’s used by the inner function even after `counter()` returns.

---

## Checking Escape Behavior

You can use this command to check all allocations:
```bash
go build -gcflags="-m" main.go
```
or
```bash
go run -gcflags="-m" main.go
```

It shows which variables were moved to the heap.

---

## Summary

| Concept | Explanation |
|----------|--------------|
| **Escape analysis** | Compiler check to decide if variable should go on stack or heap |
| **Escapes when** | Variable or its address is used outside its function |
| **Stack allocation** | Fast, temporary storage within a function |
| **Heap allocation** | Slower, persistent storage managed by garbage collector |
| **How to check** | Use `go run -gcflags="-m"` |
| **Goal** | Efficient memory use and safe pointer management |

---

## In One Line

> Escape Analysis in Go ensures safe and efficient memory allocation by deciding whether a variable stays on the stack or escapes to the heap based on how it’s used.
