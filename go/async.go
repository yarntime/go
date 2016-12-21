package main

import (
	"fmt"
	"time"
)

var ch = make(chan int)

func main() {

	for i := 0; i < 4; i++ {
		go worker(i)
	}

	for i := 0; i < 1000; i++ {
		ch <- i
	}
}

func worker(i int) {
	for {
		select {
		case num := <-ch:
			fmt.Printf("worker %d print %d\n", i, num)
			time.Sleep(time.Second)
		}
	}
}
