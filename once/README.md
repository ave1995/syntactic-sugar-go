# 🥇 sync.Once

## Purpose
A **Once** structure is used to guarantee that a function, often for initialization purposes, is executed **exactly once**, regardless of how many goroutines attempt to call it concurrently. This is commonly used for implementing **singletons**.

## Key Methods
* `Do(f func())`: Executes the function `f`. If `Do` has already been called successfully, subsequent calls (even concurrent ones) will ignore `f`.

## Best Practice
Use `sync.Once` for **lazy initialization** of expensive or necessary resources that must only be set up one time throughout the application's lifecycle, such as configuration loading or establishing a global connection pool.