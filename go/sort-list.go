package main

import "fmt"

type Node struct {
	Value int
	Next  *Node
}

func sort(list1, list2 *Node) *Node {
	var result *Node
	var tmp *Node
	if list1 == nil {
		return list2
	}

	if list2 == nil {
		return list1
	}

	if list1.Value < list2.Value {
		result = list1
		list1 = list1.Next
		tmp = result
	} else {
		result = list2
		list2 = list2.Next
		tmp = result
	}

	for list1 != nil && list2 != nil {
		if list1.Value < list2.Value {
			tmp.Next = list1
			list1 = list1.Next
		} else {
			tmp.Next = list2
			list2 = list2.Next
		}
		tmp = tmp.Next
	}

	if list1 != nil {
		tmp.Next = list1
	} else if list2 != nil {
		tmp.Next = list2
	}

	return result
}

func main() {
	list1 := &Node{
		Value: 1,
		Next: &Node{
			Value: 3,
			Next:  nil,
		},
	}

	list2 := &Node{
		Value: 2,
		Next: &Node{
			Value: 4,
			Next:  nil,
		},
	}

	list3 := sort(list1, list2)
	for list3 != nil {
		fmt.Printf("%d->", list3.Value)
		list3 = list3.Next
	}
	fmt.Println()
}
