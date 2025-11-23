package main

import "sync"

func main() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			// ... work logic ...
		}(i)
	}

	wg.Wait() // Blocks until all 5 workers call Done()
}
