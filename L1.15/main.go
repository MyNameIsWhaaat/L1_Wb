package main

import (
	"fmt"
)

var _ = fmt.Println

func createHugeString(n int) string { return "" }

func firstNBytes(s string, n int) string {
	if n >= len(s) {
		return s
	}
	b := make([]byte, n)
	copy(b, s[:n])
	return string(b)
}

func someFunc() string {
	v := createHugeString(1 << 20)
	return firstNBytes(v, 100)
}

func main() {
	justString := someFunc()
	fmt.Println(len(justString))
}