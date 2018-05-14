package main

import (
	"math"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(4)
	for i := 0; i < 1000000; i++ {
		go test()
	}
	select {}
}

func test() {
	for i := 0; i < 100000000; i++ {
		for j := 0; j < 100000000; j++ {
			for k := 0; k < 100000000; k++ {
				a := 4 * math.Pow(-1, float64(i)) * math.Pow(-1, float64(j)) * math.Pow(-1, float64(k))
				math.Sin(a)
				math.Sincos(a)
				math.Atan2(a, math.Atanh(a))
			}
		}
	}
}
