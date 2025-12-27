package main

import (
	"fmt"
	"time"
)

func rangeOverSlice() {
	fmt.Println("\n=== 1. RANGE OVER SLICE ===")

	numbers := []int{10, 20, 30, 40, 50}

	// Range returns: index, value
	fmt.Println("With index and value:")
	for index, value := range numbers {
		fmt.Printf("  Index: %d, Value: %d\n", index, value)
	}

	// Only values (ignore index with _)
	fmt.Println("\nOnly values:")
	for _, value := range numbers {
		fmt.Printf("  Value: %d\n", value)
	}

	// Only indexes
	fmt.Println("\nOnly indexes:")
	for index := range numbers {
		fmt.Printf("  Index: %d\n", index)
	}
}

func rangeOverArray() {
	fmt.Println("\n=== 2. RANGE OVER ARRAY ===")

	colors := [4]string{"red", "green", "blue", "yellow"}

	for i, color := range colors {
		fmt.Printf("  colors[%d] = %s\n", i, color)
	}
}

func rangeOverMap() {
	fmt.Println("\n=== 3. RANGE OVER MAP ===")

	ages := map[string]int{
		"Alice":   25,
		"Bob":     30,
		"Charlie": 35,
	}

	// Range returns: key, value
	fmt.Println("Keys and values:")
	for name, age := range ages {
		fmt.Printf("  %s is %d years old\n", name, age)
	}

	// Only keys
	fmt.Println("\nOnly keys:")
	for name := range ages {
		fmt.Printf("  %s\n", name)
	}

	// NOTE: Map iteration order is NOT guaranteed!
	fmt.Println("\nNote: Order may vary on different runs")
}

func rangeOverString() {
	fmt.Println("\n=== 4. RANGE OVER STRING(rune) ===")

	text := "Hello, 世界"

	// Range over string returns: index, rune (unicode code point)
	fmt.Println("Character by character:")
	for index, runeValue := range text {
		fmt.Printf("  Index: %d, Rune: %c (Unicode: %U)\n", index, runeValue, runeValue)
	}

	// Note: Index jumps for multi-byte characters!
	fmt.Println("\nNotice: Index isn't sequential for multi-byte chars (世界)")
}

func rangeOverChannel() {
	fmt.Println("\n=== 5. RANGE OVER CHANNEL ===")

	ch := make(chan int, 5)

	// Send data to channel
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i * 10
			time.Sleep(200 * time.Millisecond)
		}
		close(ch) // IMPORTANT: Must close channel for range to end!
	}()

	// Range over channel reads until channel is closed
	fmt.Println("Reading from channel:")
	for value := range ch {
		fmt.Printf("  Received: %d\n", value)
	}

	fmt.Println("Channel closed, range loop ended")
}

func rangeWithBreakContinue() {
	fmt.Println("\n=== 6. RANGE WITH BREAK AND CONTINUE ===")

	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Println("Skip even numbers, stop at 8:")
	for _, num := range numbers {
		if num%2 == 0 {
			continue // Skip even numbers
		}
		if num > 8 {
			break // Stop when we reach a number > 8
		}
		fmt.Printf("  %d\n", num)
	}
}

func rangeOverNested() {
	fmt.Println("\n=== 7. RANGE OVER NESTED STRUCTURES ===")

	matrix := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println("2D Matrix:")
	for i, row := range matrix {
		for j, value := range row {
			fmt.Printf("  matrix[%d][%d] = %d\n", i, j, value)
		}
	}
}

func rangeOverStructSlice() {
	fmt.Println("\n=== 8. RANGE OVER STRUCT SLICE ===")

	type Person struct {
		Name string
		Age  int
	}

	people := []Person{
		{Name: "Alice", Age: 25},
		{Name: "Bob", Age: 30},
		{Name: "Charlie", Age: 35},
	}

	for i, person := range people {
		fmt.Printf("  Person %d: %s, age %d\n", i+1, person.Name, person.Age)
	}

	// Important: range creates a COPY of each element
	fmt.Println("\nModifying via range (won't work):")
	for _, person := range people {
		person.Age += 1 // This modifies the COPY, not original!
	}
	fmt.Printf("  Alice's age is still: %d (unchanged)\n", people[0].Age)

	fmt.Println("\nModifying correctly using index:")
	for i := range people {
		people[i].Age += 1 // This modifies the original
	}
	fmt.Printf("  Alice's age is now: %d (changed)\n", people[0].Age)
}

func rangeOverMapOfSlices() {
	fmt.Println("\n=== 9. RANGE OVER MAP OF SLICES ===")

	teams := map[string][]string{
		"Backend":  {"Alice", "Bob"},
		"Frontend": {"Charlie", "David"},
		"DevOps":   {"Eve"},
	}

	for teamName, members := range teams {
		fmt.Printf("Team %s:\n", teamName)
		for i, member := range members {
			fmt.Printf("  %d. %s\n", i+1, member)
		}
	}
}

func rangePerformanceNote() {
	fmt.Println("\n=== 10. RANGE PERFORMANCE NOTE ===")

	// Large slice
	data := make([]int, 5)
	for i := range data {
		data[i] = i * 100
	}

	// Method 1: Range with index and value (creates copy)
	fmt.Println("Method 1 - with value (copy created):")
	for i, value := range data {
		fmt.Printf("  [%d]=%d ", i, value)
	}

	// Method 2: Range with only index (no copy)
	fmt.Println("\n\nMethod 2 - index only (more efficient):")
	for i := range data {
		fmt.Printf("  [%d]=%d ", i, data[i])
	}

	fmt.Println("\n\nNote: For large structs/slices, use index-only for better performance")
}

// ============================================
// MAIN - run all examples
// ============================================
func main() {
	fmt.Println("🔁 GO RANGE - COMPLETE GUIDE")
	fmt.Println("==============================")

	rangeOverSlice()
	time.Sleep(300 * time.Millisecond)

	rangeOverArray()
	time.Sleep(300 * time.Millisecond)

	rangeOverMap()
	time.Sleep(300 * time.Millisecond)

	rangeOverString()
	time.Sleep(300 * time.Millisecond)

	rangeOverChannel()
	time.Sleep(300 * time.Millisecond)

	rangeWithBreakContinue()
	time.Sleep(300 * time.Millisecond)

	rangeOverNested()
	time.Sleep(300 * time.Millisecond)

	rangeOverStructSlice()
	time.Sleep(300 * time.Millisecond)

	rangeOverMapOfSlices()
	time.Sleep(300 * time.Millisecond)

	rangePerformanceNote()

	fmt.Println("\n\n✅ All examples completed!")
}
