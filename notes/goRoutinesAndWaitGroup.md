
# Goroutines, WaitGroups, and Scheduler in Go

This guide explains **Goroutines**, **WaitGroups**, and the **Go Scheduler** — in a simple, clear, and easy-to-understand way. It also covers common interview questions so you can confidently discuss concurrency in Go.

---

##  1. What Are Goroutines?

- A **goroutine** is a lightweight thread managed by the Go runtime (not by the OS).
- They allow Go to perform **concurrent** tasks efficiently.
- Goroutines start with the `go` keyword.

Example:

```go
go task(i)
```

This tells Go: “Run `task(i)` concurrently (in the background).”

---

###  Key Points

- Each goroutine starts with only **~2 KB** of memory.
- You can create **thousands or millions** of them.
- The main function **does not wait** for goroutines to finish by default.
- If the main function exits early, **all running goroutines stop immediately**.

---

###  Example

```go
for i := 0; i <= 10; i++ {
    go task(i)
}
time.Sleep(time.Second) // prevents main from exiting too early
```

- `go task(i)` runs tasks concurrently.
- `time.Sleep()` keeps the main function alive temporarily.
- This is not ideal because we don’t know how long the tasks will take.

---

##  2. WaitGroups — Proper Way to Wait for Goroutines

`sync.WaitGroup` is used to **synchronize** goroutines.  
It ensures all goroutines finish before the program continues.

### Example

```go
var wg sync.WaitGroup

for i := 0; i <= 10; i++ {
    wg.Add(1)            // increment counter
    go task(i, &wg)      // start goroutine
}

wg.Wait()                // wait until counter == 0
```

Inside each goroutine:

```go
func task(id int, wg *sync.WaitGroup) {
    defer wg.Done()      // decrement counter when finished
    fmt.Println("Task", id)
}
```

---

###  How WaitGroup Works

1. `wg.Add(1)` — increases an internal counter by 1.  
   Do this **before** starting the goroutine.
2. `wg.Done()` — decreases the counter by 1.
3. `wg.Wait()` — blocks until counter = 0 (all done).

The counter goes up and down like this:

```
Add → Add → Add → Done → Done → ... → counter = 0 → main continues
```

---

###  Add/Done Timing Clarification

The loop **does not fill the counter first**.  
Goroutines may start running while the loop continues.

For example:

```
i=0: Add(+1) → start G0
G0 runs fast → Done(-1)
i=1: Add(+1) → start G1
```
So `Add` and `Done` can interleave — the scheduler decides.  
That’s why `Add(1)` must always happen **before** starting each goroutine.

---

###  Defer in WaitGroups

- `defer wg.Done()` ensures `Done()` runs when the goroutine finishes.
- Even if the goroutine panics or returns early, `Done()` will still run.
- Works like a “cleanup” or destructor.

---

##  3. The Go Scheduler and Multiplexing

When you run 1,000 goroutines on a CPU with 4 cores — how does Go handle it?

Go’s **scheduler** maps **many goroutines (N)** onto **fewer OS threads (M)** efficiently.  
This is called **M:N scheduling**.

### ⚙️ Example Visualization

| OS Threads | Goroutines (running one at a time per thread) |
|-------------|-----------------------------------------------|
| Thread 1    | G1, G4, G7, G10                              |
| Thread 2    | G2, G5, G8                                   |
| Thread 3    | G3, G6, G9                                   |

The scheduler quickly switches goroutines across OS threads — this is **multiplexing**.

You can think of it as **many goroutines sharing few threads**.

---

###  Internal Scheduler Components

| Term | Meaning | Description |
|------|----------|-------------|
| **M (Machine)** | OS thread | Actual OS-level thread |
| **P (Processor)** | Logical processor | Holds a queue of runnable goroutines |
| **G (Goroutine)** | Lightweight function | The actual goroutine unit |

Each **P** is linked to one **M**, and holds many **G**s.  
The scheduler assigns, pauses, and resumes goroutines as needed.

If one M is idle, Go performs **work stealing** — it takes runnable goroutines from another M to balance load.

---

###  Why It’s Powerful

- Allows millions of goroutines without memory issues.
- Efficiently uses all CPU cores.
- You never have to manage threads manually.
- Fair scheduling and fast switching are handled by Go itself.

---

###  Summary of Key Concepts

| Concept | Meaning |
|----------|----------|
| **Goroutine** | Lightweight concurrent function |
| **WaitGroup** | Synchronization mechanism to wait for completion |
| **Add(1)** | Increments internal counter |
| **Done()** | Decrements internal counter |
| **Wait()** | Blocks until counter = 0 |
| **Scheduler** | Decides when and where goroutines run |
| **Multiplexing** | Many goroutines share few OS threads |
| **M:N Model** | Maps many goroutines (N) onto fewer OS threads (M) |
| **Defer** | Ensures code runs after function exits |

---


## 4. Common Interview Questions About Goroutines & WaitGroups

This file contains clear, simple explanations for common questions about **Goroutines**, **WaitGroups**, and the **Go Scheduler**. It’s written in plain language for easy understanding and interview preparation.

---

## 1. What exactly is a goroutine? How is it different from a thread?

A **goroutine** is a lightweight thread managed by the **Go runtime**, not by the operating system.

**OS threads** are heavy — they take more memory (1–2 MB each) and are slower to create and switch between.

**Goroutines**, on the other hand, are extremely lightweight — each starts with about **2 KB of stack space**, which automatically grows and shrinks as needed.

The Go runtime scheduler manages goroutines efficiently by **multiplexing many goroutines onto a smaller number of OS threads**. This is called **M:N scheduling**.

You can run thousands or even millions of goroutines at once without running out of memory.

**Key idea:** The Go runtime scheduler manages *many goroutines (N)* across *few OS threads (M)*.

---

## 2. How does Go achieve concurrency?

Go achieves **concurrency** using **goroutines** and **channels**.

- **Concurrency** means doing many things *at once in structure* (they make progress together, not necessarily simultaneously).
- **Parallelism** means *executing* multiple things at the *same time* (on multiple CPU cores).

The Go runtime’s scheduler decides which goroutines run in parallel based on available CPU cores.

**Simple rule:**  
Concurrency is about *design* — parallelism is about *execution*.

---

## 3. What is the role of sync.WaitGroup in goroutines?

A **WaitGroup** helps **synchronize** multiple goroutines.

It ensures that the **main goroutine waits** until all other goroutines finish executing.

Internally, it uses a **counter** that tracks active goroutines.

- `Add(1)` increases the counter.
- `Done()` decreases the counter.
- `Wait()` blocks until the counter becomes zero.

This guarantees that your program only continues once all goroutines have completed.

---

## 4. What happens if you forget to call wg.Done()?

If a goroutine never calls `wg.Done()`, the internal counter never reaches zero.

That means `wg.Wait()` will **block forever**, and the program will hang.

This situation is called a **deadlock** — when goroutines are waiting forever because the counter is unbalanced.

**Key term:** Deadlock — when goroutines wait forever due to missing `Done()` calls.

---

## 5. What happens if you call wg.Done() more times than Add()?

If you call `Done()` more times than `Add()`, the counter becomes negative.

Go will panic with this error:

```
panic: sync: negative WaitGroup counter
```

This means some goroutine called `Done()` without a matching `Add()`.

**Key idea:** Negative counter happens due to a mismatch between `Add` and `Done`.

---

## 6. Why do we use defer wg.Done() instead of wg.Done() directly?

`defer wg.Done()` ensures that `Done()` runs automatically when the goroutine finishes — even if it returns early or panics.

This is safer and avoids forgetting to decrease the counter manually.

However, `defer` is slightly slower in performance-critical code, so sometimes developers call `wg.Done()` manually.

**Example:**

```go
func task(id int, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Println("Task:", id)
}
```

---

## 7. Why do we call Add(1) before starting a goroutine?

Because if the goroutine starts and finishes *before* `Add(1)` executes, it can cause a **race condition** where the counter goes negative.

To prevent that, you should always call `Add(1)` **before** the goroutine starts.

**Correct:**

```go
wg.Add(1)
go task(&wg)
```

**Incorrect:**

```go
go task(&wg)
wg.Add(1) // unsafe
```

---

## 8. What is a race condition in goroutines?

A **race condition** happens when two or more goroutines access the same variable at the same time, and at least one of them modifies it, without proper synchronization.

This causes unpredictable results.

You can detect race conditions using Go’s built-in race detector:

```bash
go run -race main.go
```

**Key idea:** Data races occur when shared memory is accessed concurrently without synchronization.

---

## 9. What are channels, and how do they differ from WaitGroups?

A **WaitGroup** is only used to wait for goroutines to finish.

A **channel** is used for **communication** between goroutines — sending and receiving data safely.

They also help with synchronization and follow Go’s philosophy:

> “Don’t communicate by sharing memory; share memory by communicating.”

So:  
- Use **WaitGroups** when you just need to wait for completion.  
- Use **channels** when goroutines need to send/receive data.

---

## 10. Can you cause a deadlock with WaitGroups?

Yes. Here’s a simple example:

```go
wg.Add(1)
wg.Wait() // main goroutine waits here forever
go func() {
    wg.Done()
}()
```

In this example, `wg.Wait()` blocks the main goroutine before the new goroutine starts, so the counter never reaches zero.

**Lesson:** Always start goroutines *before* calling `Wait()`.

---

## 11. What is the difference between concurrency and parallelism?

- **Concurrency:** Multiple tasks make progress at the same time (not necessarily simultaneously).  
- **Parallelism:** Multiple tasks execute at the same time on multiple CPUs.

Goroutines provide concurrency, and if multiple CPU cores are available, the Go runtime may run them in parallel.

---

## 12. What is GOMAXPROCS?

`GOMAXPROCS` defines how many CPU cores Go can use to run goroutines in parallel.

You can check or set it using:

```go
runtime.GOMAXPROCS(4)
```

By default, it equals the number of available CPU cores on your system.

**Key term:** GOMAXPROCS controls the number of OS threads that can execute Go code simultaneously.

---

## 13. What happens if the main function exits while goroutines are still running?

When the main function finishes, the entire program exits immediately — even if goroutines are still running.

All active goroutines are stopped abruptly.

That’s why we use **WaitGroups** or **channels** to make sure all background goroutines complete before the program ends.

---

## 14. How does Go’s scheduler manage goroutines?

Go’s scheduler is part of the Go runtime. It uses an **M:N model** to map many goroutines (N) onto fewer OS threads (M).

The scheduler decides:

- Which goroutine runs next
- Which OS thread runs it
- When to pause or resume goroutines

If one thread becomes idle, Go uses **work stealing** to balance goroutines across threads.

This automatic management is what makes goroutines so efficient.

**Key term:** M:N scheduler — many goroutines mapped to fewer OS threads efficiently.

---

## 15. Can you run millions of goroutines?

Yes. Go is designed for high-concurrency systems.

Each goroutine uses only a few kilobytes of memory and dynamically adjusts its stack size as needed.

That’s why Go can handle **hundreds of thousands** or even **millions of concurrent goroutines**, making it ideal for scalable servers and microservices.

---

## Summary

| Concept | Explanation |
|----------|-------------|
| Goroutine | Lightweight concurrent function managed by Go runtime |
| WaitGroup | Used to synchronize multiple goroutines |
| Add / Done / Wait | Counter mechanism for tracking goroutines |
| Defer | Ensures cleanup executes when function ends |
| Scheduler | Maps many goroutines to fewer OS threads |
| M:N Model | Many goroutines run across fewer OS threads |
| Race Condition | Unsynchronized access to shared data |
| Deadlock | Happens when goroutines wait forever |
| GOMAXPROCS | Defines CPU cores available for parallel execution |

---

### Final Definition

Goroutines are lightweight, concurrent functions managed by Go’s runtime. The Go scheduler maps many goroutines (N) onto fewer OS threads (M) to achieve efficient concurrency. WaitGroups are used to ensure that all goroutines complete before the program exits.

---

##  Bonus Concepts

### Race Condition Example

```go
count := 0
for i := 0; i < 5; i++ {
    go func() { count++ }()
}
fmt.Println(count)
```
This can print unpredictable results. Use `sync.Mutex` or channels to fix it.

---

### Scheduler Analogy

Imagine a restaurant:
- **Waiters (OS threads)** handle multiple **customers (goroutines)**.
- The **manager (scheduler)** decides which waiter serves which customers.
- Customers come and go — waiters are reused instead of hiring new ones.

That’s how goroutines share OS threads efficiently.

---

## Final Summary

| Concept | Description |
|----------|-------------|
| **Goroutine** | Lightweight concurrent task managed by Go runtime |
| **WaitGroup** | Used to wait for multiple goroutines to finish |
| **Add/Done/Wait** | Counter-based synchronization system |
| **Defer** | Ensures cleanup runs at end of function |
| **Scheduler** | Manages which goroutine runs where |
| **Multiplexing** | Many goroutines share fewer OS threads |
| **M:N Model** | Many goroutines mapped to few OS threads |
| **GOMAXPROCS** | Limits CPU cores used for parallel execution |
| **Deadlock** | Happens when main waits forever (missing Done) |
| **Race Condition** | Happens when shared data accessed unsafely |

---

### One-Liner Definition

> **Goroutines are lightweight, concurrent functions managed by Go’s runtime. The Go scheduler maps many goroutines (N) onto fewer OS threads (M) to achieve concurrency efficiently. WaitGroups are used to ensure all goroutines complete before exiting the program.**
