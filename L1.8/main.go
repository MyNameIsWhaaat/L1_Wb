package main

import (
	"fmt"
)

func SetBit(n int64, i uint, val int) int64 {
	if val == 0 {
		return n &^ (1 << i)
	}
	return n | (1 << i)
}

func main() {
	var num int64 = 5

	fmt.Printf("num = %d (%04b)\n", num, num)

	res := SetBit(num, 1, 0)
	fmt.Printf("SetBit(num, 1, 0) = %d (%04b)\n", res, res)

	res = SetBit(num, 2, 1)
	fmt.Printf("SetBit(num, 2, 1) = %d (%04b)\n", res, res)
}