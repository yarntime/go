package main

import (
	"fmt"
)

func count(x int) int {
	result := 0
	for x != 0 {
		result++
		x = x & (x - 1)
	}
	return result
}

func main() {
	fmt.Printf("%d\n", count(999))
}
