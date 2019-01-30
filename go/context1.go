package main

import (
	"fmt"
	"time"
	"context"
)

func Concurrency(ctx context.Context, sn int) {
	interval := time.After(1 * time.Second)
	cur := 0
	for {
		select {
		case <- interval:
			fmt.Printf("%d print: %d\n", sn, cur)
			cur++
			interval = time.After(1 * time.Second)
		case <- ctx.Done():
			fmt.Printf("exit.\n")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithCancel(context.TODO())
	for i := 0; i < 3; i++ {
		go Concurrency(ctx, i)
	}

	time.Sleep(10 * time.Second)
	cancel()
	time.Sleep(1 * time.Second)
}
