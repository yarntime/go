package main

import (
	"fmt"
	"sync"
	"time"
)

var ch = make(chan int)
var finishCount int
var l sync.Mutex

func main() {

	finishCount = 0

	for i := 0; i < 4; i++ {
		go worker(i)
	}

	for i := 0; i < 100000; i++ {
		ch <- i
	}

	time.Sleep(time.Second)
	fmt.Printf("finish count is %d\n", finishCount)
	close(ch)
}

func worker(i int) {
	for {
		select {
		case num := <-ch:
			fmt.Printf("worker %d print %d\n", i, num)
			time.Sleep(time.Nanosecond * 1000)
			l.Lock()
			finishCount++
			l.Unlock()
		}
	}
}
