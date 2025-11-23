# 📦 sync.Pool

## Purpose
A **Pool** provides a temporary collection of objects that can be reused by multiple goroutines. Its primary goal is to **reduce memory allocation and garbage collection overhead** by recycling objects that are expensive to create, like large I/O buffers.

## Key Methods
* `New func() interface{}`: A field in the `Pool` struct. This function is called to create a new object when `Get()` is called and the pool is empty.
* `Get() interface{}`: Retrieves a cached object from the pool. If the pool is empty, it calls the `New` function.
* `Put(x interface{})`: Returns an object `x` to the pool for reuse.

## Important Note
Pooled objects are subject to garbage collection, meaning the objects are **not guaranteed to persist** in the pool across different GC cycles. Do not use `sync.Pool` for persistent storage or stateful objects.