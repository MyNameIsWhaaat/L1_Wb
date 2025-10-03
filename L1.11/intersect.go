package main

func Intersect(a, b []int) []int {
	set := make(map[int]struct{})
	var result []int

	for _, v := range a {
		set[v] = struct{}{}
	}

	for _, v := range b {
		if _, ok := set[v]; ok {
			result = append(result, v)
			delete(set, v)
		}
	}

	return result
}
