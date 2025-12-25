package main

import "fmt"

func LetterFrequencies(word string) {
	m := make(map[rune]int, len(word))

	for _, v := range word {
		m[v]++
	}

	for key, v := range m {
		fmt.Printf("%c has count %d.\n", key, v)
	}
}

func main() {
	LetterFrequencies("aabbc")
}
