package main

import (
	"fmt"
	"time"
)

func test() {
	chexit := make(chan int)
	chforever := make(chan int)
	go exit(chexit)
	go forever(chforever)
	select {
	case <-chexit:
		fmt.Println("finished.")
		chforever <- 1
	}
}

func exit(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Printf("exit %d\n", i)
		time.Sleep(1 * time.Second)
	}
	ch <- 1
}

func forever(ch chan int) {
	var i = 0
	for {
		select {
		case <-ch:
			fmt.Println("forever exit.")
			return
		default:
			fmt.Printf("forever %d\n", i)
			i++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	test()
	fmt.Println("main process exit.")
}
