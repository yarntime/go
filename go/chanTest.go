package main

import (
	"fmt"
)

func main() {
	printNum := make(chan bool)
	printChar := make(chan bool)
	exit := make(chan bool, 2)
	go func() {
		for i := 1; i < 11; i += 2 {
			fmt.Printf("%d", i)
			fmt.Printf("%d", i + 1)
			printChar <- true
			<- printNum
		}
		exit <- true
	}()

	go func() {
		for i := 1; i < 11; i += 2 {
			<- printChar
			fmt.Printf("%c", i - 1 + 'A')
			fmt.Printf("%c", i + 'A')
			printNum <- true
		}
		exit <- true
	}()

	<- exit
	<- exit
}