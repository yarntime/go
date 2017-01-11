package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("start")
	chn := make(chan int, 5)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		x := rand.Intn(5)
		fmt.Println("i is", i, "rand is:", x)
		go worker(i, x, chn)

	}
	fmt.Println("end")
	for i := 0; i < 5; i++ {
		j := <-chn
		fmt.Println(j)
	}

}

func worker(i int, sleepInt int, chn chan int) {
	d := time.Duration(sleepInt) * time.Second
	time.Sleep(d)
	fmt.Println("I am ", i)
	j := i + 10
	chn <- j
}