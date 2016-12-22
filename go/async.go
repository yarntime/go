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

	for {
		fmt.Printf("finish count is %d\n", finishCount)
		if finishCount == 100000 {
			break
		}
		time.Sleep(time.Second)
	}
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
