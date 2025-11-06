package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	r := []rune(s)
	for i, j := 0, len(r)-1; i < j; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	in := bufio.NewReader(os.Stdin)
	text, _ := in.ReadString('\n')
	if n := len(text); n > 0 && (text[n-1] == '\n' || text[n-1] == '\r') {
		text = text[:n-1]
	}
	fmt.Println(reverseString(text))
}
