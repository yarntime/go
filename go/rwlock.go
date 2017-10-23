package main

import (
	"sync"
	"time"
)

var m *sync.RWMutex
var w sync.WaitGroup

func main() {
	m = new(sync.RWMutex)
	go write(1)
	w.Add(1)
	go read(2)
	w.Add(1)
	go write(3)
	w.Add(1)
	w.Wait()
}

func read(i int) {
	println(i, "read start")
	m.RLock()
	defer m.RUnlock()
	defer w.Done()
	println(i, "reading")
	time.Sleep(1 * time.Second)
	println(i, "read end")
}

func write(i int) {
	println(i, "write start")
	m.Lock()
	defer m.Unlock()
	defer w.Done()
	println(i, "writing")
	time.Sleep(1 * time.Second)
	println(i, "write end")
}
