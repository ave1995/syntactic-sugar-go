package main

import (
	"fmt"
	"time"
)

func basicWhileLoop() {
	fmt.Println("\n=== 1. BASIC 'WHILE' LOOP ===")

	counter := 0

	// This is like: while (counter < 5)
	for counter < 5 {
		fmt.Printf("Counter: %d\n", counter)
		counter++
	}

	fmt.Println("Loop finished!")
}

func infiniteLoop() {
	fmt.Println("\n=== 2. INFINITE LOOP ===")

	counter := 0

	// This is like: while (true)
	for {
		fmt.Printf("Iteration: %d\n", counter)
		counter++

		// Break condition
		if counter >= 5 {
			fmt.Println("Breaking out of infinite loop!")
			break
		}

		time.Sleep(300 * time.Millisecond)
	}
}

func whileWithContinue() {
	fmt.Println("\n=== 3. WHILE WITH CONTINUE ===")

	num := 0

	for num < 10 {
		num++

		// Skip even numbers
		if num%2 == 0 {
			continue
		}

		fmt.Printf("Odd number: %d\n", num)
	}
}

func whileMultipleConditions() {
	fmt.Println("\n=== 4. WHILE WITH MULTIPLE CONDITIONS ===")

	x := 0
	y := 10

	// Loop while x < 5 AND y > 5
	for x < 5 && y > 5 {
		fmt.Printf("x=%d, y=%d\n", x, y)
		x++
		y--
	}

	fmt.Println("Condition no longer met!")
}

func doWhilePattern() {
	fmt.Println("\n=== 5. DO-WHILE PATTERN ===")

	counter := 0

	// Execute at least once, then check condition
	for {
		fmt.Printf("Executed: %d\n", counter)
		counter++

		// Check condition at the end (like do-while)
		if !(counter < 3) {
			break
		}
	}
}

func whileWithChannels() {
	fmt.Println("\n=== 6. WHILE WITH CHANNELS ===")

	ch := make(chan int, 5)

	// Send some data
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(200 * time.Millisecond)
		}
		close(ch)
	}()

	// Read while channel is open
	received := 0
	for received < 5 {
		value := <-ch
		fmt.Printf("Received: %d\n", value)
		received++
	}
}

func nestedWhileLoops() {
	fmt.Println("\n=== 7. NESTED WHILE LOOPS ===")

	i := 0
	for i < 3 {
		j := 0
		for j < 3 {
			fmt.Printf("i=%d, j=%d\n", i, j)
			j++
		}
		i++
	}
}

func whileWithLabel() {
	fmt.Println("\n=== 8. WHILE WITH LABEL ===")

	i := 0

OuterLoop:
	for i < 5 {
		j := 0
		for j < 5 {
			fmt.Printf("i=%d, j=%d\n", i, j)

			// Break outer loop when i=2 and j=2
			if i == 2 && j == 2 {
				fmt.Println("Breaking outer loop!")
				break OuterLoop
			}
			j++
		}
		i++
	}
}

func traditionalForLoop() {
	fmt.Println("\n=== 9. TRADITIONAL FOR LOOP ===")

	// Classic for loop with init, condition, post
	for i := 0; i < 5; i++ {
		fmt.Printf("Traditional for: i=%d\n", i)
	}
}

func retryPattern() {
	fmt.Println("\n=== 10. RETRY PATTERN ===")

	maxRetries := 3
	attempts := 0
	success := false

	for attempts < maxRetries && !success {
		attempts++
		fmt.Printf("Attempt %d of %d...\n", attempts, maxRetries)

		// Simulate operation that might fail
		if attempts == 2 {
			success = true
			fmt.Println("Success!")
		} else {
			fmt.Println("Failed, retrying...")
			time.Sleep(500 * time.Millisecond)
		}
	}

	if !success {
		fmt.Println("All retries exhausted!")
	}
}

func main() {
	fmt.Println("🔄 GO WHILE LOOPS - COMPLETE GUIDE")
	fmt.Println("==================================")

	basicWhileLoop()
	time.Sleep(300 * time.Millisecond)

	infiniteLoop()
	time.Sleep(300 * time.Millisecond)

	whileWithContinue()
	time.Sleep(300 * time.Millisecond)

	whileMultipleConditions()
	time.Sleep(300 * time.Millisecond)

	doWhilePattern()
	time.Sleep(300 * time.Millisecond)

	whileWithChannels()
	time.Sleep(300 * time.Millisecond)

	nestedWhileLoops()
	time.Sleep(300 * time.Millisecond)

	whileWithLabel()
	time.Sleep(300 * time.Millisecond)

	traditionalForLoop()
	time.Sleep(300 * time.Millisecond)

	retryPattern()

	fmt.Println("\n✅ All examples completed!")
}
