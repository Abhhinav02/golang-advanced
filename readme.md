# 🌱 What is a Goroutine?

A **goroutine** is a lightweight, independently executing function that runs **concurrently** with other goroutines in the same address space.
Think of it as:

* In **JavaScript**, we have an **event loop** that handles async tasks (e.g., promises, async/await).
* In **Go**, instead of a single-threaded event loop, we have **goroutines managed by the Go runtime**.

They allow us to perform tasks like handling requests, I/O operations, or computations in parallel **without manually managing threads**.

---

# ⚖️ Goroutine vs OS Thread

| Feature                   | Goroutine                        | OS Thread               |
| ------------------------- | -------------------------------- | ----------------------- |
| **Size at start**         | ~2 KB stack                      | ~1 MB stack             |
| **Managed by**            | Go runtime scheduler (M:N model) | OS Kernel               |
| **Number you can create** | Millions                         | Limited (few thousands) |
| **Switching**             | Very fast, done in user space    | Slower, done by OS      |
| **Creation cost**         | Extremely cheap                  | Expensive               |

👉 This is why we say goroutines are *lightweight threads*.

---

# ⚙️ How to Start a Goroutine

```go
package main

import (
	"fmt"
	"time"
)

func printMessage(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go printMessage("goroutine") // runs concurrently
	printMessage("main")         // runs in main goroutine
}
```

* The `go` keyword starts a new goroutine.
* Here:

  * `main()` itself runs in the **main goroutine**.
  * `go printMessage("goroutine")` starts another goroutine.
* If `main()` exits before the new goroutine finishes, the program ends immediately.

⚠️ Unlike JavaScript promises (which keep the process alive until settled), Go doesn’t wait for goroutines unless you **explicitly synchronize** them.

---

# 🧵 Go’s Concurrency Model (M:N Scheduler)

Go runtime uses an **M:N scheduler**, meaning:

* **M goroutines** are multiplexed onto **N OS threads**.
* This is different from **1:1** (like Java threads) or **N:1** (like cooperative multitasking).

The scheduler ensures:

* Goroutines are distributed across multiple threads.
* When one blocks (e.g., waiting on I/O), another is scheduled.

Think of goroutines as **tasks in a work-stealing scheduler**.

---

# 🛠️ Synchronization with Goroutines

Since goroutines run concurrently, we need synchronization tools:

### 1. **WaitGroup** – Wait for Goroutines to Finish

```go
package main

import (
	"fmt"
	"sync"
)

func worker(id int, wg *sync.WaitGroup) {
	defer wg.Done() // signals completion
	fmt.Printf("Worker %d starting\n", id)
	// simulate work
	fmt.Printf("Worker %d done\n", id)
}

func main() {
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)            // add to wait counter
		go worker(i, &wg)
	}

	wg.Wait() // wait for all to finish
}
```

✅ Ensures the program won’t exit before all goroutines finish.

---

### 2. **Channels** – Communication Between Goroutines

Channels are **Go’s big idea** for concurrency.
Instead of sharing memory and locking it, goroutines **communicate by passing messages**.

```go
package main

import "fmt"

func worker(ch chan string) {
	ch <- "task finished" // send data into channel
}

func main() {
	ch := make(chan string)

	go worker(ch)

	msg := <-ch // receive data
	fmt.Println("Message:", msg)
}
```

👉 Think of it like JavaScript `Promise.resolve("task finished")`, but **synchronous communication** unless buffered.

---

### 3. **Buffered Channels** – Queue of Messages

```go
ch := make(chan int, 2) // capacity = 2
ch <- 10
ch <- 20
fmt.Println(<-ch)
fmt.Println(<-ch)
```

* Unbuffered channel: send blocks until receive is ready.
* Buffered channel: send doesn’t block until buffer is full.

---

### 4. **select** – Multiplexing Channels

```go
select {
case msg := <-ch1:
	fmt.Println("Received", msg)
case msg := <-ch2:
	fmt.Println("Received", msg)
default:
	fmt.Println("No message")
}
```

Like `Promise.race()` in JS.

---

# 🔥 Key Gotchas with Goroutines

1. **Main goroutine exit kills all child goroutines**.
   → Always use WaitGroups or channels to synchronize.

2. **Race conditions** happen if goroutines write/read shared data without sync.
   → Use `sync.Mutex`, `sync.RWMutex`, or better: **channels**.

3. **Too many goroutines** can cause memory pressure, but still far cheaper than threads.

4. **Don’t block forever** – unreceived channel sends cause deadlocks.

---

# 📊 Real-World Use Cases

* **Web servers**: Each request can run in its own goroutine.
* **Scraping / Crawling**: Launch a goroutine for each URL fetch.
* **Background jobs**: Run tasks concurrently (DB writes, logging, metrics).
* **Pipelines**: Process data in multiple stages with goroutines + channels.

---

# 🧠 Mental Model (JS vs Go)

* **JavaScript** → concurrency = single-threaded event loop + async callbacks.
* **Go** → concurrency = many goroutines scheduled onto multiple OS threads.

So:

* In JS, concurrency = illusion via async.
* In Go, concurrency = real, parallel execution when multiple CPU cores exist.

---

✅ To summarize:

* Goroutines = **cheap concurrent tasks** managed by Go runtime.
* Not OS threads, but multiplexed onto threads.
* Communicate via **channels** instead of shared memory.
* Powerful with **WaitGroups, select, and synchronization tools**.

---

**concurrency vs parallelism** is a core concept in computer science and in Go (since Go was built with concurrency in mind). Let’s break it down step by step in detail.

---

## **1. The Core Idea**

* **Concurrency** = Dealing with many tasks at once (managing multiple things).
* **Parallelism** = Doing many tasks at the same time (executing multiple things simultaneously).

Both sound similar, but they’re not the same.

---

## **2. Analogy**

Imagine we’re in a restaurant kitchen:

* **Concurrency (chef multitasking):**
  One chef handles multiple dishes by switching between them. He cuts vegetables for Dish A, stirs the sauce for Dish B, and checks the oven for Dish C. He’s *not doing them at the exact same time*, but he’s managing multiple tasks *in progress*.

* **Parallelism (many chefs working together):**
  Three chefs cook three different dishes at the *same time*. Tasks truly happen *simultaneously*.

👉 Concurrency is about **structure** (how tasks are managed).
👉 Parallelism is about **execution** (how tasks are run in hardware).

---

## **3. Technical Definition**

* **Concurrency**:
  Multiple tasks *make progress* in overlapping time periods. It doesn’t require multiple processors/cores. Even with a single CPU core, the system can *interleave execution* of tasks via context switching.

* **Parallelism**:
  Multiple tasks *run at the exact same instant*, usually on different CPU cores or processors.

---

## **4. Example with Go**

Go is famous for concurrency with **goroutines**.

```go
package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(500 * time.Millisecond)
	}
}

func main() {
	go task("Task A") // run concurrently
	go task("Task B")

	time.Sleep(3 * time.Second)
	fmt.Println("Done")
}
```

### What happens:

* **Concurrency:** Both `Task A` and `Task B` *appear to run at the same time* because Go schedules goroutines across available cores. If you run this on a single-core CPU, Go interleaves execution → that’s concurrency.
* **Parallelism:** If you run this on a multi-core CPU, `Task A` might run on Core 1 and `Task B` on Core 2 simultaneously → that’s parallelism.

---

## **5. Key Differences Table**

| Aspect              | Concurrency                         | Parallelism                                    |
| ------------------- | ----------------------------------- | ---------------------------------------------- |
| **Definition**      | Managing multiple tasks at once     | Executing multiple tasks at once               |
| **Focus**           | Task switching and scheduling       | Simultaneous execution                         |
| **CPU Requirement** | Can happen on a single-core CPU     | Requires multi-core CPU                        |
| **Analogy**         | One chef multitasking across dishes | Many chefs cooking different dishes            |
| **In Go**           | Achieved via goroutines & channels  | Achieved when goroutines run on multiple cores |

---

## **6. Visual Representation**

* **Concurrency (single-core):**

```
Time: |----A----|----B----|----A----|----B----|
       ^ Task A and Task B interleaved
```

* **Parallelism (multi-core):**

```
Core1: |----A----|----A----|----A----|
Core2: |----B----|----B----|----B----|
       ^ Tasks running truly at the same time
```

---

## **7. In Practice**

* Concurrency is **a design approach**: "How do we structure a program so that it can handle many things at once?"
* Parallelism is **an execution strategy**: "How do we use hardware to literally do many things at once?"

Go is *concurrent by design* (goroutines + channels) and *parallel by runtime* (GOMAXPROCS decides how many cores are used).

---

✅ **Final takeaway**:

* **Concurrency = composition of independently executing tasks.**
* **Parallelism = simultaneous execution of tasks.**

They are related, but not the same. A program can be concurrent but not parallel, parallel but not concurrent, or both.

---






