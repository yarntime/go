package main

import "fmt"

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func getPos(arr []int, s, e int) int {
	i, j := s+1, e
	for i < j {
		for arr[j] >= arr[s] {
			j--
		}
		for arr[i] <= arr[s] {
			i++
		}
		if i < j {
			swap(&arr[i], &arr[j])
		}

	}
	swap(&arr[s], &arr[j])
	return j
}

func qsort(arr []int, s, e int) {
	if s >= e {
		return
	}
	pos := getPos(arr, s, e)
	qsort(arr, s, pos-1)
	qsort(arr, pos+e, e)
}

func main() {
	p := []int{4, 1, 2, 6, 4, 3}
	qsort(p, 0, 5)
	for i := 0; i < len(p); i++ {
		fmt.Print(p[i], " ")
	}
}
