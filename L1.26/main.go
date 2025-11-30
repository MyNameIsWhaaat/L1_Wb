package main

import (
	"fmt"
	"unicode"
)

func HasAllUniqueChars(s string) bool {
	seen := make(map[rune]struct{})

	for _, r := range s {
		lr := unicode.ToLower(r)

		if _, exists := seen[lr]; exists {
			return false
		}
		seen[lr] = struct{}{}
	}

	return true
}

func main() {
	tests := []string{
		"abcd",
		"abCdefAaf",
		"aabcd",
		"Привет",
		"ПривЕтп",
	}

	for _, t := range tests {
		fmt.Printf("%q -> %v\n", t, HasAllUniqueChars(t))
	}
}
