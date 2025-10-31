package main

import "fmt"

func partition(arr []int, low, high int) int {

	pivot := arr[(low+high)/2]

	i := low - 1
	for j := low; j <= high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}

	return i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		pivotIndex := partition(arr, low, high)
		
		quickSort(arr, low, pivotIndex)
		quickSort(arr, pivotIndex+1, high)
	}
	return arr
}

func main() {

	arr := []int{10, 7, 8, 9, 1, 5}
	fmt.Println("Original array:", arr)

	sortedArr := quickSort(arr, 0, len(arr)-1)

	fmt.Println("Sorted array:", sortedArr)
}
