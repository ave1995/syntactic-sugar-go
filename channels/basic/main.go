package main

import (
	"fmt"
	"time"
)

func unbufferedChannel() {
	fmt.Println("\n=== 1. UNBUFFERED CHANNEL ===")

	ch := make(chan int)

	// Goroutine that sends data to the channel
	go func() {
		fmt.Println("Goroutine: Sending number 42 to channel...")
		ch <- 42 // WRITE - sends data to channel (blocks until someone reads)
		fmt.Println("Goroutine: Number sent!")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("Main: Read from channel...")
	number := <-ch // READ - reads data from channel (blocks until data available)
	fmt.Printf("Main: Got number: %d\n", number)
}

func bufferedChannel() {
	fmt.Println("\n=== 2. BUFFERED CHANNEL ===")

	// Create a channel with buffer for 3 elements
	ch := make(chan string, 3)

	// We can send 3 values without blocking!
	ch <- "first"
	ch <- "second"
	ch <- "third"
	fmt.Println("Sent 3 values to buffered channel (without block)")

	// Reading values
	fmt.Println("Read:", <-ch)
	fmt.Println("Read:", <-ch)
	fmt.Println("Read:", <-ch)

	// Attempting to send 4th value without goroutine would cause deadlock
}

func goroutineCommunication() {
	fmt.Println("\n=== 3. COMMUNICATION BETWEEN GOROUTINES ===")

	ch := make(chan int)

	// Worker goroutine
	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Printf("Worker: Sending %d\n", i)
			ch <- i
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// Main reads 5 values
	for i := 1; i <= 5; i++ {
		value := <-ch
		fmt.Printf("Main: Received %d\n", value)
	}
}

func multipleGoroutines() {
	fmt.Println("\n=== 4. MULTIPLE GOROUTINES WITH CHANNEL ===")

	ch := make(chan string, 5)

	// 3 worker goroutines
	for i := 1; i <= 3; i++ {
		workerID := i
		go func() {
			message := fmt.Sprintf("Message from worker %d", workerID)
			ch <- message
		}()
	}

	// Read 3 messages
	for i := 1; i <= 3; i++ {
		message := <-ch
		fmt.Println("Received:", message)
	}
}

func blockingDemo() {
	fmt.Println("\n=== 5. BLOCKING DEMONSTRATION ===")

	ch := make(chan int)

	go func() {
		fmt.Println("Goroutine: Waiting 2 seconds before sending...")
		time.Sleep(2 * time.Second)
		fmt.Println("Goroutine: Sending data!")
		ch <- 100
	}()

	fmt.Println("Main: Waiting for data from channel (blocking)...")
	data := <-ch
	fmt.Printf("Main: Finally received: %d\n", data)
}

func channelTypes() {
	fmt.Println("\n=== 6. CHANNEL TYPES ===")

	// Channel for int
	intCh := make(chan int)

	// Channel for string
	strCh := make(chan string, 2)

	// Channel for struct
	type Message struct {
		ID   int
		Text string
	}
	msgCh := make(chan Message, 1)

	// Usage
	go func() {
		intCh <- 42
		strCh <- "Hello"
		msgCh <- Message{ID: 1, Text: "Test message"}
	}()

	fmt.Printf("Int channel: %d\n", <-intCh)
	fmt.Printf("String channel: %s\n", <-strCh)
	msg := <-msgCh
	fmt.Printf("Struct channel: ID=%d, Text=%s\n", msg.ID, msg.Text)
}

func main() {
	fmt.Println("GO CHANNELS - Complete guide")
	fmt.Println("============================")

	unbufferedChannel()
	time.Sleep(500 * time.Millisecond)

	bufferedChannel()
	time.Sleep(500 * time.Millisecond)

	goroutineCommunication()
	time.Sleep(500 * time.Millisecond)

	multipleGoroutines()
	time.Sleep(500 * time.Millisecond)

	blockingDemo()
	time.Sleep(500 * time.Millisecond)

	channelTypes()
}
