package main

import "fmt"

type Element interface{}

type Queue interface {
	Offer(e Element)
	Poll() Element
	Clear() bool
	Size() int
	IsEmpty() bool
}

type sliceEntry struct {
	element []Element
}

func NewQueue() *sliceEntry {
	return &sliceEntry{}
}

func (entry *sliceEntry) Offer(e Element) {
	entry.element = append(entry.element, e)
}

func (entry *sliceEntry) Poll() Element {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return nil
	}

	firstElement := entry.element[0]
	entry.element = entry.element[1:]
	return firstElement
}

func (entry *sliceEntry) Clear() bool {
	if entry.IsEmpty() {
		fmt.Println("queue is empty!")
		return false
	}
	for i := 0; i < entry.Size(); i++ {
		entry.element[i] = nil
	}
	entry.element = nil
	return true
}

func (entry *sliceEntry) Size() int {
	return len(entry.element)
}

func (entry *sliceEntry) IsEmpty() bool {
	if len(entry.element) == 0 {
		return true
	}
	return false
}

type Node struct {
	Value       Element
	Left, Right *Node
}

func PrintTree(root *Node) {
	queue := NewQueue()
	queue.Offer(root)
	curRight := root
	tmpRight := root

	for !queue.IsEmpty() {
		node := queue.Poll().(*Node)
		fmt.Printf("%v", node.Value)

		if node.Left != nil {
			queue.Offer(node.Left)
			tmpRight = node.Left
		}

		if node.Right != nil {
			queue.Offer(node.Right)
			tmpRight = node.Right
		}

		if node == curRight {
			curRight = tmpRight
			fmt.Println()
		} else {
			fmt.Printf(",")
		}
	}
}

func main() {
	tree := &Node{
		Value: 2,
		Left: &Node{
			Value: 4,
			Left: &Node{
				Value: 5,
				Left:  nil,
				Right: &Node{
					Value: 3,
					Left:  nil,
					Right: nil,
				},
			},
			Right: &Node{
				Value: 8,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &Node{
			Value: 10,
			Left:  nil,
			Right: nil,
		},
	}
	PrintTree(tree)
}
