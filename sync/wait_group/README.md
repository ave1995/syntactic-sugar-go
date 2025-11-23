# ⏳ sync.WaitGroup

## Purpose
A **WaitGroup** is used to wait for a collection of goroutines to finish executing. It allows the main goroutine to block until all worker goroutines have completed their tasks.

## Key Methods
* `Add(delta int)`: Increments the internal counter by `delta`. This is called *before* launching the goroutine.
* `Done()`: Decrements the internal counter by one. This is typically called using `defer` inside the worker goroutine.
* `Wait()`: Blocks the calling goroutine until the internal counter is zero.

## Workflow
1.  Call `Add(N)` in the main goroutine (where N is the number of workers).
2.  Launch the N worker goroutines.
3.  Call `defer wg.Done()` inside each worker.
4.  Call `Wait()` in the main goroutine to block and wait.