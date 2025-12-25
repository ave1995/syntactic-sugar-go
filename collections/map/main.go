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

func WhatHappend() {
	var m map[string]int

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from panic:", r)
		}
	}()

	val := m["key"] // safe
	fmt.Printf("I read from nil map:  %d\n", val)

	val, ok := m["key"]
	fmt.Printf("I check and read from nil map: %d (ok=%v)\n", val, ok)

	m["key"] = 1
}

func main() {
	LetterFrequencies("aabbc")

	WhatHappend()
}
