package main

import (
	"fmt"
)

func reverseWords(s string) string {
	r := []rune(s)

	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	start := 0
	for i := 0; i <= len(r); i++ {
		if i == len(r) || r[i] == ' ' {
			for l, h := start, i-1; l < h; l, h = l+1, h-1 {
				r[l], r[h] = r[h], r[l]
			}
			start = i + 1
		}
	}
	return string(r)
}

func main() {
	fmt.Println(reverseWords("snow dog sun"))
}
