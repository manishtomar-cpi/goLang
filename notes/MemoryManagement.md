# Go Memory Management Explained

Go manages memory automatically, meaning you don't need to manually allocate or free memory (like in C or C++).  
It uses a garbage collector (GC) and decides where data should be stored — either on the stack or the heap.

---

## 1. Stack vs Heap

### Stack
- The stack is used for local variables inside functions.
- Memory allocation and deallocation are fast (Last In, First Out — LIFO).
- Automatically freed when a function returns.
- No need for garbage collection here.

#### Example:
```go
func add(a, b int) int {
    sum := a + b   // 'sum' lives on the stack
    return sum
}
```
Here, `sum` is a local variable. It exists only while `add()` runs, and memory is released when the function ends.

---

### Heap
- The heap is used for variables that outlive the function call.
- Accessed when data needs to be shared across functions or goroutines.
- Managed by the garbage collector.
- Allocation is slower than stack allocation.

#### Example:
```go
func createPointer() *int {
    x := 10
    return &x // 'x' escapes to the heap
}
```
Here, Go places `x` on the heap because we return its address. It must remain valid after the function exits.

---

## 2. Escape Analysis

Go’s compiler performs something called escape analysis to decide:
- If a variable stays on the stack (safe, short-lived)
- Or needs to move to the heap (longer-lived)

You can check escape analysis using:
```
go build -gcflags="-m" main.go
```
It will show messages like:
```
moved to heap: x
```

---

## 3. Garbage Collection (GC)

Go uses automatic garbage collection to clean up heap memory.  
The garbage collector runs in the background and removes data that:
- Is no longer reachable (no active references)
- Is not in use by any part of the program

### Key Points:
- Go’s GC is concurrent and non-blocking.
- It runs automatically — you don’t manually trigger it.
- It minimizes "stop-the-world" pauses for smooth performance.
- Optimized for low latency and high throughput.

---

## 4. How GC Works (Simplified)

1. Mark Phase → Go marks all objects still in use (reachable).  
2. Sweep Phase → It frees memory used by unreachable objects.  
3. Compact Phase → Adjusts memory if needed (to prevent fragmentation).

All of this happens automatically.

---

## 5. Summary Table

| Concept | Description | Managed By | Speed |
|----------|--------------|-------------|--------|
| Stack | Local, short-lived variables | Automatically (no GC) | Very fast |
| Heap | Long-lived, shared variables | Garbage Collector | Slower |
| Escape Analysis | Decides stack vs heap | Compiler | - |
| Garbage Collection | Frees unused heap memory | Runtime | Automatic |

---

## 6. Best Practices
- Avoid unnecessary pointers (they can force heap allocation).
- Reuse memory when possible (e.g., with slices or sync.Pool).
- Keep functions small so most variables stay on the stack.
- Let the Go runtime manage memory — you focus on logic.

---

## 7. Example of Both
```go
package main

import "fmt"

func stackExample() {
    a := 5
    b := 10
    fmt.Println(a + b) // 'a' and 'b' are on the stack
}

func heapExample() *int {
    x := 42
    return &x // 'x' goes to heap because we return its address
}

func main() {
    stackExample()
    ptr := heapExample()
    fmt.Println(*ptr)
}
```
Here:
- `a` and `b` → allocated on stack  
- `x` → allocated on heap  
- `ptr` → points to heap memory managed by the garbage collector

---

# In Short
- Stack → fast, temporary memory  
- Heap → long-term, garbage-collected memory  
- Go runtime automatically decides what goes where and when to free it.
