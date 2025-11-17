package main

func DeleteAtKeepOrder[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	copy(s[i:], s[i+1:])
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}

func DeleteAtSwapLast[T any](s []T, i int) []T {
	if i < 0 || i >= len(s) {
		return s
	}
	s[i] = s[len(s)-1]
	var zero T
	s[len(s)-1] = zero
	return s[:len(s)-1]
}

func main() {
	s := []string{"a", "b", "c", "d"}
	s = DeleteAtKeepOrder(s, 1)

	t := []int{10, 20, 30, 40}
	t = DeleteAtSwapLast(t, 1)
}
