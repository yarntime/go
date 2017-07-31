package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan int)
	ticker := time.NewTicker(5 * time.Second)
	defer func() {
		ticker.Stop()
	}()

	for {
		select {
		case <-c:

		case <-ticker.C:
			fmt.Println("triggled...")
		}
	}
}
