package main

import "fmt"

func echoArray(a interface{}) {
	b, _ := a.([]int)
	for _, v := range b {
		fmt.Print(v, " ")
	}
	fmt.Println()
	return
}

func main() {
	a := []int{2, 1, 3, 5, 4}
	echoArray(a)
}
