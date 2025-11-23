# 🔒 sync.Mutex

## Purpose
A **Mutex** (Mutual Exclusion Lock) is used to protect shared data from concurrent access by multiple goroutines. It ensures that only one goroutine can access a critical section of code at any given time, preventing **data races**.

## Key Methods
* `Lock()`: Acquires the lock. If another goroutine holds the lock, the caller blocks until it's released.
* `Unlock()`: Releases the lock. Must be called when the access to the shared resource is complete.

## Best Practice
Always use the `defer` keyword with `Unlock()` immediately after calling `Lock()`. This guarantees the lock is released, even if the function panics or returns early.