package main

import (
	"fmt"
	"time"
)

func basicPanic() {
	fmt.Println("\n=== 1. BASIC PANIC ===")

	fmt.Println("Before panic")
	panic("Something went wrong!") // Program crashes here
	fmt.Println("This will NEVER print")
}

func panicWithDifferentTypes() {
	fmt.Println("\n=== 2. PANIC WITH DIFFERENT TYPES ===")

	// Can panic with any type
	// panic("string message")
	// panic(42)
	// panic([]string{"error", "list"})

	type CustomError struct {
		Code    int
		Message string
	}

	panic(CustomError{Code: 500, Message: "Server error"})
}

func deferWithPanic() {
	fmt.Println("\n=== 3. DEFER WITH PANIC ===")

	defer fmt.Println("Defer 1: This runs even if panic happens!")
	defer fmt.Println("Defer 2: Defers run in LIFO order (Last In First Out)")

	fmt.Println("Before panic")
	panic("Panic occurred!")
	fmt.Println("This won't print")
}

func recoverFromPanic() {
	fmt.Println("\n=== 4. RECOVER FROM PANIC ===")

	// Defer with recover to catch panic
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic: %v\n", r)
		}
	}()

	fmt.Println("About to panic...")
	panic("This panic will be recovered!")
	fmt.Println("This won't print")
}

func recoverOutsideDefer() {
	fmt.Println("\n=== 5. RECOVER ONLY WORKS IN DEFER ===")

	// This won't work - recover must be in defer!
	if r := recover(); r != nil {
		fmt.Println("This won't catch anything")
	}

	// Correct way
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Correctly recovered: %v\n", r)
		}
	}()

	panic("Testing recover placement")
}

func multipleDeferWithPanic() {
	fmt.Println("\n=== 6. MULTIPLE DEFERS WITH PANIC ===")

	defer fmt.Println("Defer 3 (executes first - LIFO)")
	defer fmt.Println("Defer 2 (executes second)")
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Defer 1 (executes third): Recovered from: %v\n", r)
		}
	}()

	fmt.Println("Normal execution")
	panic("Panic in multiple defer example")
}

func panicInGoroutine() {
	fmt.Println("\n=== 7. PANIC IN GOROUTINE ===")

	// Launch goroutine with panic handling
	go func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Goroutine recovered from: %v\n", r)
			}
		}()

		fmt.Println("Goroutine: About to panic...")
		panic("Panic in goroutine!")
	}()

	time.Sleep(500 * time.Millisecond)
	fmt.Println("Main: Continues running after goroutine panic")
}

func panicWithoutRecover() {
	fmt.Println("\n=== 8. PANIC WITHOUT RECOVER (COMMENTED) ===")
	fmt.Println("If we uncomment the panic, program will crash")

	// Uncomment to see the crash:
	// panic("Unhandled panic - program will crash!")

	fmt.Println("Program continues...")
}

func safeDivision(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("panic occurred: %v", r)
			result = 0
		}
	}()

	result = a / b // Will panic if b is 0
	return result, nil
}

func divisionExample() {
	fmt.Println("\n=== 9. REAL-WORLD EXAMPLE - SAFE DIVISION ===")

	// Safe division
	result, err := safeDivision(10, 2)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 2 = %d\n", result)
	}

	// Division by zero (will panic but we catch it)
	result, err = safeDivision(10, 0)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("10 / 0 = %d\n", result)
	}

	fmt.Println("Program continues after recovered panic")
}

func nestedPanicRecover() {
	fmt.Println("\n=== 10. NESTED PANIC AND RECOVER ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Outer recover caught: %v\n", r)
		}
	}()

	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("Inner recover caught: %v\n", r)
				// Can panic again!
				panic("Re-panicking from inner function")
			}
		}()

		panic("Initial panic")
	}()
}

func outOfBoundsPanic() {
	fmt.Println("\n=== 11. OUT OF BOUNDS PANIC ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from out of bounds: %v\n", r)
		}
	}()

	slice := []int{1, 2, 3}
	fmt.Printf("Slice: %v\n", slice)
	fmt.Println("Trying to access index 10...")

	// This will panic with "index out of range"
	_ = slice[10]
}

func nilPointerPanic() {
	fmt.Println("\n=== 12. NIL POINTER PANIC ===")

	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from nil pointer: %v\n", r)
		}
	}()

	type Person struct {
		Name string
	}

	var p *Person // nil pointer
	fmt.Println("Trying to access nil pointer...")

	// This will panic with "nil pointer dereference"
	fmt.Println(p.Name)
}

func panicVsError() {
	fmt.Println("\n=== 13. PANIC VS ERROR ===")

	fmt.Println("USE ERRORS (normal approach):")
	fmt.Println("  - Expected failures")
	fmt.Println("  - User input validation")
	fmt.Println("  - Network errors")
	fmt.Println("  - File not found")

	fmt.Println("\nUSE PANIC (exceptional situations):")
	fmt.Println("  - Programming errors")
	fmt.Println("  - Impossible situations")
	fmt.Println("  - Initialization failures")
	fmt.Println("  - Critical invariants violated")

	fmt.Println("\nExample: Use error for file operations")
	// Good: return error
	// Bad: panic("file not found")
}

func main() {
	fmt.Println("💥 GO PANIC AND RECOVER - COMPLETE GUIDE")
	fmt.Println("==========================================")

	// Note: We wrap dangerous functions in their own calls
	// so they can demonstrate panic/recover without crashing

	// Example 1: Would crash without recover
	func() {
		defer func() {
			if r := recover(); r != nil {
				fmt.Printf("\nCaught panic from basicPanic: %v\n", r)
			}
		}()
		basicPanic()
	}()

	time.Sleep(300 * time.Millisecond)

	// Example 3: Defer with panic
	func() {
		defer func() {
			recover() // Catch to continue
		}()
		deferWithPanic()
	}()

	time.Sleep(300 * time.Millisecond)

	recoverFromPanic()
	time.Sleep(300 * time.Millisecond)

	recoverOutsideDefer()
	time.Sleep(300 * time.Millisecond)

	multipleDeferWithPanic()
	time.Sleep(300 * time.Millisecond)

	panicInGoroutine()
	time.Sleep(700 * time.Millisecond)

	panicWithoutRecover()
	time.Sleep(300 * time.Millisecond)

	divisionExample()
	time.Sleep(300 * time.Millisecond)

	nestedPanicRecover()
	time.Sleep(300 * time.Millisecond)

	outOfBoundsPanic()
	time.Sleep(300 * time.Millisecond)

	nilPointerPanic()
	time.Sleep(300 * time.Millisecond)

	panicVsError()

	fmt.Println("\n✅ All examples completed without crashing!")
}
