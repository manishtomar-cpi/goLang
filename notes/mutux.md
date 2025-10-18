
# Mutex in Go – Complete Explanation

A **Mutex** (short for *Mutual Exclusion*) is one of the most important synchronization tools in Go. It prevents **race conditions** by allowing only one goroutine at a time to access shared data. Without proper synchronization, multiple goroutines can modify a shared variable simultaneously, resulting in unpredictable or incorrect results.

---

## 1. Why Mutex Is Needed

In concurrent programs, multiple goroutines may try to read or modify the same variable at the same time.  
When that happens without coordination, the goroutines can **overwrite each other’s work**, causing incorrect outcomes.

### Example: Without Mutex (Race Condition)

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

type post struct {
    view int
}

func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()
    p.view++
}

func main() {
    var wg sync.WaitGroup
    myPost := post{view: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg)
    }

    wg.Wait()
    fmt.Println("Total views:", myPost.view)
}
```

**Expected output:**  
```
Total views: 100
```

**Actual output (often):**  
```
Total views: 95
```
or any random number less than 100.

This happens because multiple goroutines are modifying `p.view` simultaneously. Each one reads the old value, increments it, and writes it back — but these operations overlap and interfere with each other.

This is called a **race condition**.

---

## 2. What Is a Mutex

A **mutex** is a locking mechanism that ensures **only one goroutine can execute a section of code at a time**.  
The Go standard library provides this in the `sync` package.

When a goroutine locks a mutex using `Lock()`, no other goroutine can lock it again until it is unlocked using `Unlock()`.

### Real-Life Analogy

Imagine there is only **one key to a room**:

- When a person (goroutine) enters the room, they **lock the door** (`Lock()`).
- No one else can enter until they **unlock the door** (`Unlock()`).
- This ensures that only one person can be in the room at any time.

---

## 3. Using Mutex in Go

Here’s the same example, now using a mutex to fix the race condition:

```go
package main

import (
    "fmt"
    "sync"
)

type post struct {
    view int
    mu   sync.Mutex
}

func (p *post) inc(wg *sync.WaitGroup) {
    defer wg.Done()

    p.mu.Lock()   // Lock before modifying shared data
    p.view++      // Only one goroutine can execute this line at a time
    p.mu.Unlock() // Unlock after modification is complete
}

func main() {
    var wg sync.WaitGroup
    myPost := post{view: 0}

    for i := 0; i < 100; i++ {
        wg.Add(1)
        go myPost.inc(&wg)
    }

    wg.Wait()
    fmt.Println("Total views:", myPost.view)
}
```

Now the output will always be:
```
Total views: 100
```

---

## 4. How Mutex Works Internally

- When `Lock()` is called:
  - If the mutex is **unlocked**, it becomes **locked** and the goroutine continues.
  - If it’s **already locked**, the goroutine is **blocked** (paused) until another goroutine calls `Unlock()`.

- When `Unlock()` is called:
  - The mutex becomes **available** again.
  - One of the blocked goroutines is allowed to continue and acquire the lock.

This ensures **only one goroutine** is inside the protected (critical) section at any given time.

---

## 5. Why Use `defer` for Unlock

You often see mutexes used with `defer`, like this:

```go
func (p *post) inc(wg *sync.WaitGroup) {
    p.mu.Lock()
    defer func() {
        p.mu.Unlock()
        wg.Done()
    }()
    p.view++
}
```

Using `defer` ensures that even if:
- The function returns early, or  
- A panic occurs,

the mutex will still be **unlocked properly**.  

If a goroutine forgets to unlock the mutex, other goroutines will wait forever — causing a **deadlock**.

---

## 6. Best Practice – Lock Only What’s Necessary

Locking a large block of code unnecessarily can reduce concurrency and slow down the program.

###  Bad Practice:
```go
p.mu.Lock()
fmt.Println("Processing...")
time.Sleep(1 * time.Second)
p.view++
p.mu.Unlock()
```
Here, the mutex is locked during printing and sleeping — which doesn’t need protection.  
This means other goroutines must wait longer than necessary.

###  Good Practice:
```go
fmt.Println("Processing...")
time.Sleep(1 * time.Second)
p.mu.Lock()
p.view++
p.mu.Unlock()
```
Now, only the shared variable update (`p.view++`) is protected, allowing maximum concurrency.

---

## 7. Deadlocks and Common Mistakes

A **deadlock** occurs when a mutex is locked and never unlocked.  
This can happen if:
- You forget to call `Unlock()`.
- You use multiple mutexes incorrectly (e.g., locking them in the wrong order).

Example:
```go
p.mu.Lock()
// some logic that panics
return // forgot to unlock!
```

This permanently locks the mutex — no other goroutine can proceed.

Using `defer` prevents this issue.

---

## 8. Mutex vs RWMutex

Go also provides a **Read-Write Mutex (`sync.RWMutex`)**.  
It allows:
- **Multiple readers** to access shared data at the same time.
- **Only one writer** to modify data (blocks all readers).

This is more efficient when you have many reads but few writes.

### Example:
```go
type post struct {
    view int
    mu   sync.RWMutex
}

func (p *post) readView() int {
    p.mu.RLock()
    defer p.mu.RUnlock()
    return p.view
}

func (p *post) inc() {
    p.mu.Lock()
    p.view++
    p.mu.Unlock()
}
```

Use `RLock()` and `RUnlock()` for reading, `Lock()` and `Unlock()` for writing.

---

## 9. When to Use Mutex

| Situation | Use Mutex? |
|------------|------------|
| Multiple goroutines modify shared variable |  Yes |
| Only one goroutine accesses variable |  No need |
| Goroutines only read shared variable |  Use RWMutex |
| Data passed via channels |  Not required — channels handle synchronization |

---

## 10. Key Points to Remember

| Concept | Description |
|----------|-------------|
| **Mutex** | Prevents multiple goroutines from accessing shared data simultaneously |
| **Lock()** | Acquire the lock — blocks if already locked |
| **Unlock()** | Release the lock so others can proceed |
| **Critical Section** | Code between Lock() and Unlock() |
| **Defer Unlock()** | Ensures unlock always happens |
| **Deadlock** | When a mutex is locked and never unlocked |
| **RWMutex** | Allows multiple readers or one writer |

---

## 11. Summary

- A **mutex** ensures that shared data is accessed safely by only one goroutine at a time.  
- Always **Lock** before modifying shared data and **Unlock** after.
- Use `defer` to guarantee unlock even during errors.
- Keep the **critical section small** to maintain performance.
- For read-heavy workloads, prefer **RWMutex**.
- Forgetting to unlock leads to **deadlocks**; not locking leads to **race conditions**.

---

> **In short:**  
> A `sync.Mutex` acts as a simple lock that ensures only one goroutine at a time can modify a shared resource.  
> It prevents race conditions, ensures data safety, and keeps concurrent code predictable and stable.
