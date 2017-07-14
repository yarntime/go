package main

import (
	"fmt"
)

type Interface interface {
	Stop()
}

type Bmw struct {
	Name string
}

func (b *Bmw) Stop() {
	fmt.Printf("%s bmw stop.", b.Name)
}

type Car struct {
	Interface
}

func main() {
	c := Car{
		&Bmw{
			Name: "i8",
		},
	}
	c.Stop()
}
