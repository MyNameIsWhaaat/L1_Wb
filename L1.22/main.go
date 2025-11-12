package main

import (
	"fmt"
	"math/big"
)

func main() {
	a := new(big.Int)
	b := new(big.Int)

	a.SetString("1234567890123456789012345", 10)
	b.SetString("9876543210987654321", 10)

	sum := new(big.Int).Add(a, b)
	diff := new(big.Int).Sub(a, b)
	prod := new(big.Int).Mul(a, b)
	quot := new(big.Int).Div(a, b)

	fmt.Println("a + b =", sum)
	fmt.Println("a - b =", diff)
	fmt.Println("a * b =", prod)
	fmt.Println("a / b =", quot)

	ar := new(big.Rat).SetInt(a)
	br := new(big.Rat).SetInt(b)
	divExact := new(big.Rat).Quo(ar, br)
	fmt.Println("a / b (exact) =", divExact.FloatString(20))
}