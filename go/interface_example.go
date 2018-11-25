package main

import (
	"fmt"
)

type Role interface {
	Do()
}

type Teacher struct {
}

func (t *Teacher) Do() {
	fmt.Printf("I'm a teacher\n")
}

type Student struct {
}

func (s *Student) Do() {
	fmt.Printf("I'm a student\n")
}

var _ Role = new(Teacher)

var _ Role = new(Student)

func Do(v interface{}) {
	t, ok := v.(Role)
	if !ok {
		fmt.Printf("type mismatched.\n")
		return
	}
	t.Do()
}

func main() {
	t := &Teacher{}
	Do(t)

	s := &Student{}
	Do(s)
}