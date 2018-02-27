package main

import (
	"fmt"
)

type student struct {
	Name string
	Age  int
}

func main() {
	st := []*student{
		{
			Name: "zhang",
			Age:  15,
		},
		{
			Name: "zhao",
			Age:  16,
		},
	}

	for _, s := range st {
		fmt.Printf("%v\n", s)
		s.Name = "test"
	}

	for _, s := range st {
		fmt.Printf("%v\n", s)
	}
}
