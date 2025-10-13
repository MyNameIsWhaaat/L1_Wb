package main

import "fmt"

func main() {
	a, b := 5, 7

	a = a + b
	b = a - b
	a = a - b

	fmt.Println("a =", a, "b =", b)
}

