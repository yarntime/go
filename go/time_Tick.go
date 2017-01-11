package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("the time")
	tc := time.Tick(time.Second)
	for i := 1; i <= 200; i++ {
		<-tc
		fmt.Println("hello")
	}
}
