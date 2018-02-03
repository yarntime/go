package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

func NewPipe() (*os.File, *os.File, error) {
	read, write, err := os.Pipe()
	if err != nil {
		return nil, nil, err
	}
	return read, write, nil
}

func main() {

	fmt.Println("main starting")
	stop := make(chan int)
	r, w, _ := NewPipe()
	go func() {
		fmt.Println("int go run")
		msg, _ := ioutil.ReadAll(r)
		fmt.Printf("int go run get msg %s \n", string(msg))
		<-stop
	}()

	w.WriteString("hello world")
	w.Close()
	time.Sleep(time.Second * 3)
	stop <- 1
	fmt.Println("main end")

}
