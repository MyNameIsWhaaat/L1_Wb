package main

import "fmt"

func binarySearchRecursive(arr []int, target int, low int, high int) int {
	if low > high {
		return -1
	}

	mid := low + (high-low)/2

	if arr[mid] == target {
		return mid
	}

	if arr[mid] < target {
		return binarySearchRecursive(arr, target, mid+1, high)
	} else {
		return binarySearchRecursive(arr, target, low, mid-1)
	}
}

func main() {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15, 17}
	target := 9

	index := binarySearchRecursive(arr, target, 0, len(arr)-1)
	if index != -1 {
		fmt.Printf("Элемент %d найден на индексе %d\n", target, index)
	} else {
		fmt.Println("Элемент не найден")
	}
}