type Node struct {
    Val int
    Min int
}
type MinStack struct {
    Stack []Node
}


/** initialize your data structure here. */
func Constructor() MinStack {
    return MinStack{
        Stack: make([]Node, 0),
    }
}


func (this *MinStack) Push(x int)  {
    if len(this.Stack) == 0 {
        this.Stack = append(this.Stack, Node{Val: x, Min: x})
    } else {
        min := min(x, this.GetMin())
        this.Stack = append(this.Stack, Node{Val: x, Min: min})
    }
}


func (this *MinStack) Pop()  {
    this.Stack = this.Stack[0:len(this.Stack) - 1]
}


func (this *MinStack) Top() int {
    return this.Stack[len(this.Stack) - 1].Val
}


func (this *MinStack) GetMin() int {
    return this.Stack[len(this.Stack) - 1].Min
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}


/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */