package main

import (
	"fmt"
)

func main() {
	ch := make(chan int, 1000)
	exit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
	}()

	go func() {
		for i := 0; i < 10; i++ {
			a, ok := <-ch

			if !ok {
				return
			}
			fmt.Println("a: ", a)
		}
		fmt.Println("close")
		close(ch)
		exit <- 1
	}()

	fmt.Println("ok")
	<-exit
}
