package main

import (
	"fmt"
	"math"
	"time"
)

func test() {
	chex := make(chan int)
	go exit(chex)
	go forever()
	select {
	case <-chex:
		fmt.Println("finished.")
		return
	}
}

func exit(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("exit %d\n", i)
		time.Sleep(5000)
	}
	ch <- 1
}

func forever() {
	for i := 0; i < math.MaxInt64; i++ {
		fmt.Printf("forever %d\n", i)
		time.Sleep(5000)
	}
}

func main() {
	go test()
	select {}
}
