package main

import (
	"fmt"
	"time"
)

func add(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

func test2() {
	ch1 := make(chan int, 10)
	go add(ch1)
	ch2 := time.After(2 * time.Second)
	for {
		select {
		case <-ch2:
			fmt.Println("timeout")
			return
		case <-ch1:
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		}
	}
}

func test3() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	ch := make(chan int, 10)
	go add(ch)
	for {
		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		case <-ch:
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		default:
		}
	}
}

func test4() {
	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()
	ch := make(chan int, 10)
	go add(ch)
	for {
		fmt.Println("I'm in.")
		select {
		case <-ticker.C:
			fmt.Println("timeout")
			return
		default:
		}
		select {
		case <-ch:
			fmt.Println("sleep one seconds ...")
			time.Sleep(1 * time.Second)
			fmt.Println("sleep one seconds end...")
		case <-ticker.C:
			fmt.Println("timeout")
			return
		}
	}
}

func main() {
	test4()
}
