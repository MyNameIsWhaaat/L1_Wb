package main

import "fmt"

func main() {
	A := []int{1, 2, 3}
	B := []int{2, 3, 4}

	inter := Intersect(A, B)
	fmt.Println("Пересечение =", inter)
}