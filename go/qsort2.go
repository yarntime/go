package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func Quicksort(s []int) {
	wg := &sync.WaitGroup{}
	wg.Add(1)
	go quicksort(s, 0, len(s)-1, wg)
	wg.Wait()
}

func quicksort(s []int, from int, to int, done *sync.WaitGroup) {

	defer done.Done()

	if from >= to {
		return
	}

	p := partition(s, from, to)
	if p > 0 {
		done.Add(1)
		go quicksort(s, from, p-1, done)
	}
	if p+1 < len(s) {
		done.Add(1)
		go quicksort(s, p+1, to, done)
	}
}

func partition(array []int, low, high int) int {
	key := array[low]
	tmpLow := low
	tmpHigh := high
	for {
		for array[tmpHigh] > key {
			tmpHigh--
		}
		for array[tmpLow] <= key && tmpLow < tmpHigh {
			tmpLow++
		}

		if tmpLow >= tmpHigh {
			break
		}

		array[tmpLow], array[tmpHigh] = array[tmpHigh], array[tmpLow]

	}
	array[tmpLow], array[low] = array[low], array[tmpLow]
	return tmpLow
}

func main() {

	ceshi := make([]int, 10000000)
	for i := 0; i < 10000000; i++ {
		ceshi[i] = rand.Intn(100)
	}
	now := time.Now()
	Quicksort(ceshi)
	fmt.Println(time.Now().Sub(now))
}
