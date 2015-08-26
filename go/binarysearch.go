package main

import (
	"fmt"
)

func binarysearch(arr []int, x int) int {
	s, e := 0, len(arr)
	for s < e {
		pos := s + (e-s)/2
		if arr[pos] == x {
			return pos + 1
		} else if arr[pos] < x {
			s = pos + 1
		} else {
			e = pos
		}
	}
	return -1
}

func main() {
	a := []int{1, 21, 31, 44, 51, 68}
	fmt.Println(binarysearch(a, 21))
}
