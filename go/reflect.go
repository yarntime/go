package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	age  int
	name string
}

func (s *Student) SetAge(i int) {
	s.age = i
}

func (s *Student) SetName(name string) {
	s.name = name
}

func (s *Student) Show() {
	fmt.Printf("student age: %d, name: %s\n", s.age, s.name)
}

func main() {
	s := &Student{
		age:  10,
		name: "test",
	}

	ref := reflect.ValueOf(&s).Elem()
	ref.MethodByName("Show").Call(nil)

	params := make([]reflect.Value, 1)
	params[0] = reflect.ValueOf(18)
	ref.MethodByName("SetAge").Call(params)

	params[0] = reflect.ValueOf("zhangw")
	ref.MethodByName("SetName").Call(params)

	ref.MethodByName("Show").Call(nil)
}
