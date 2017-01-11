package main

import (
	"fmt"
)

func main() {
	var m1, m2 int = 1, 1
	for i := 1; i < 12; i++ {
		fmt.Println(m1, m2)
		m1 += m2
		m2 += m1
	}
}
