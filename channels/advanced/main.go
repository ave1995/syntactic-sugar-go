package main

import (
	"fmt"
	"time"
)

func closingChannels() {
	fmt.Println("\n=== 1. CLOSING CHANNELS ===")

	ch := make(chan int, 3)

	// Send some values
	ch <- 1
	ch <- 2
	ch <- 3

	close(ch) // Close the channel

	// Can still read from closed channel
	fmt.Println("Reading from closed channel:")
	fmt.Println("Value:", <-ch) // 1
	fmt.Println("Value:", <-ch) // 2
	fmt.Println("Value:", <-ch) // 3

	// Reading from empty closed channel returns zero value
	val := <-ch
	fmt.Printf("Empty closed channel returns: %d\n", val) // 0

	// Check if channel is closed
	val, ok := <-ch
	if !ok {
		fmt.Println("Channel is closed (ok=false)")
	}
}

func rangeOverClosedChannel() {
	fmt.Println("\n=== 2. RANGE OVER CLOSED CHANNEL ===")

	ch := make(chan string, 3)

	// Send and close
	ch <- "first"
	ch <- "second"
	ch <- "third"
	close(ch)

	// Range automatically stops when channel is closed
	fmt.Println("Range over closed channel:")
	for value := range ch {
		fmt.Printf("  Received: %s\n", value)
	}
	fmt.Println("Range finished (channel closed)")
}

func basicSelect() {
	fmt.Println("\n=== 3. BASIC SELECT STATEMENT ===")

	ch1 := make(chan string)
	ch2 := make(chan string)

	// Send to ch1 after 1 second
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- "Message from channel 1"
	}()

	// Send to ch2 after 500ms
	go func() {
		time.Sleep(500 * time.Millisecond)
		ch2 <- "Message from channel 2"
	}()

	// Select waits for whichever channel is ready first
	select {
	case msg1 := <-ch1:
		fmt.Println("Received:", msg1)
	case msg2 := <-ch2:
		fmt.Println("Received:", msg2)
	}

	fmt.Println("Select finished (only one case executed)")
}

func selectMultipleReady() {
	fmt.Println("\n=== 4. SELECT WITH MULTIPLE READY CHANNELS ===")

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// Both channels ready
	ch1 <- "Channel 1"
	ch2 <- "Channel 2"

	// Select randomly picks one
	select {
	case msg := <-ch1:
		fmt.Println("Selected:", msg)
	case msg := <-ch2:
		fmt.Println("Selected:", msg)
	}

	fmt.Println("Note: Selection is random when multiple cases are ready")
}

func selectWithDefault() {
	fmt.Println("\n=== 5. SELECT WITH DEFAULT CASE ===")

	ch := make(chan string)

	// Non-blocking select
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("No data available, default case executed")
	}

	// Now send data
	go func() {
		ch <- "Hello"
	}()

	time.Sleep(100 * time.Millisecond)

	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("Default")
	}
}

func selectWithTimeout() {
	fmt.Println("\n=== 6. SELECT WITH TIMEOUT ===")

	ch := make(chan string)

	// Goroutine that sends after 2 seconds
	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Data arrived"
	}()

	// Wait maximum 1 second
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout! No data received in 1 second")
	}
}

func selectInLoop() {
	fmt.Println("\n=== 7. SELECT IN LOOP ===")

	ch := make(chan int, 5)

	// Producer
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
			time.Sleep(300 * time.Millisecond)
		}
		close(ch)
	}()

	// Consumer with timeout
	for {
		select {
		case val, ok := <-ch:
			if !ok {
				fmt.Println("Channel closed, exiting loop")
				return
			}
			fmt.Printf("Received: %d\n", val)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout waiting for data")
			return
		}
	}
}

func directionalChannels() {
	fmt.Println("\n=== 8. DIRECTIONAL CHANNELS ===")

	ch := make(chan string, 2)

	// Send-only channel function
	sendOnly := func(ch chan<- string) {
		ch <- "Hello"
		ch <- "World"
		// Can't read: val := <-ch  // Would cause compile error
	}

	// Receive-only channel function
	receiveOnly := func(ch <-chan string) {
		msg1 := <-ch
		msg2 := <-ch
		fmt.Printf("Received: %s %s\n", msg1, msg2)
		// Can't send: ch <- "test"  // Would cause compile error
	}

	sendOnly(ch)
	receiveOnly(ch)
}

func selectWithSend() {
	fmt.Println("\n=== 9. SELECT WITH SEND OPERATIONS ===")

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// Try to send to either channel
	select {
	case ch1 <- "Sent to ch1":
		fmt.Println("Successfully sent to ch1")
	case ch2 <- "Sent to ch2":
		fmt.Println("Successfully sent to ch2")
	default:
		fmt.Println("Both channels are full")
	}

	// Read what was sent
	select {
	case msg := <-ch1:
		fmt.Println("Read from ch1:", msg)
	case msg := <-ch2:
		fmt.Println("Read from ch2:", msg)
	}
}

func nilChannelBehavior() {
	fmt.Println("\n=== 10. NIL CHANNEL BEHAVIOR ===")

	var ch chan string // nil channel

	// Reading/writing to nil channel blocks forever
	// This would deadlock: <-ch

	// But in select with default, it's safe
	select {
	case msg := <-ch:
		fmt.Println("Received:", msg)
	default:
		fmt.Println("Nil channel blocks, default executed")
	}

	// Useful for disabling a case in select
	fmt.Println("\nDisabling channels in select:")
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 1
	ch2 <- 2

	// Disable ch1
	ch1 = nil

	select {
	case val := <-ch1:
		fmt.Println("From ch1:", val) // Won't happen (ch1 is nil)
	case val := <-ch2:
		fmt.Println("From ch2:", val) // Will happen
	}
}

func workerPoolPattern() {
	fmt.Println("\n=== 11. WORKER POOL PATTERN ===")

	jobs := make(chan int, 5)
	results := make(chan int, 5)

	// Start 3 workers
	for w := 1; w <= 3; w++ {
		workerID := w
		go func() {
			for job := range jobs {
				fmt.Printf("Worker %d processing job %d\n", workerID, job)
				time.Sleep(200 * time.Millisecond)
				results <- job * 2
			}
		}()
	}

	// Send jobs
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)

	// Collect results
	for r := 1; r <= 5; r++ {
		result := <-results
		fmt.Printf("Result: %d\n", result)
	}
}

func fanInPattern() {
	fmt.Println("\n=== 12. FAN-IN PATTERN ===")

	// Multiple producers
	producer := func(id int) <-chan int {
		ch := make(chan int)
		go func() {
			for i := 1; i <= 3; i++ {
				ch <- id*10 + i
				time.Sleep(100 * time.Millisecond)
			}
			close(ch)
		}()
		return ch
	}

	// Fan-in: merge multiple channels into one
	fanIn := func(channels ...<-chan int) <-chan int {
		out := make(chan int)

		for _, ch := range channels {
			go func() {
				for val := range ch {
					out <- val
				}
			}()
		}

		return out
	}

	ch1 := producer(1)
	ch2 := producer(2)
	ch3 := producer(3)

	merged := fanIn(ch1, ch2, ch3)

	// Read from merged channel
	timeout := time.After(1 * time.Second)
	for {
		select {
		case val := <-merged:
			fmt.Printf("Received: %d\n", val)
		case <-timeout:
			fmt.Println("Done receiving")
			return
		}
	}
}

func pingPongPattern() {
	fmt.Println("\n=== 13. PING-PONG PATTERN ===")

	ping := make(chan string)
	pong := make(chan string)

	// Ping goroutine
	go func() {
		for i := 1; i <= 3; i++ {
			msg := <-ping
			fmt.Printf("Ping received: %s\n", msg)
			time.Sleep(200 * time.Millisecond)
			pong <- "pong"
		}
	}()

	// Main acts as pong sender
	for i := 1; i <= 3; i++ {
		ping <- "ping"
		msg := <-pong
		fmt.Printf("Pong received: %s\n", msg)
	}
}

func main() {
	fmt.Println("🚀 GO ADVANCED CHANNELS - COMPLETE GUIDE")
	fmt.Println("==========================================")

	closingChannels()
	time.Sleep(300 * time.Millisecond)

	rangeOverClosedChannel()
	time.Sleep(300 * time.Millisecond)

	basicSelect()
	time.Sleep(300 * time.Millisecond)

	selectMultipleReady()
	time.Sleep(300 * time.Millisecond)

	selectWithDefault()
	time.Sleep(300 * time.Millisecond)

	selectWithTimeout()
	time.Sleep(300 * time.Millisecond)

	selectInLoop()
	time.Sleep(300 * time.Millisecond)

	directionalChannels()
	time.Sleep(300 * time.Millisecond)

	selectWithSend()
	time.Sleep(300 * time.Millisecond)

	nilChannelBehavior()
	time.Sleep(300 * time.Millisecond)

	workerPoolPattern()
	time.Sleep(500 * time.Millisecond)

	fanInPattern()
	time.Sleep(300 * time.Millisecond)

	pingPongPattern()

	fmt.Println("\n✅ All examples completed!")
}
