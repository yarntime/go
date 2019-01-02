type MyStack struct {
    data []int
}

func (this *MyStack) IsEmpty() bool {
    return len(this.data) == 0
}

func (this *MyStack) Push(x int) {
    this.data = append(this.data, x)
}

func (this *MyStack) Pop() int {
    val := this.data[len(this.data) - 1]
    this.data = this.data[:len(this.data) - 1]
    return val
}

func (this *MyStack) Top() int {
    return this.data[len(this.data) - 1]
}

type MyQueue struct {
    S1 MyStack
    S2 MyStack
}


/** Initialize your data structure here. */
func Constructor() MyQueue {
    return MyQueue{S1: MyStack{[]int{}}, S2: MyStack{[]int{}}}
}


/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int)  {
    this.S1.Push(x)
}


/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
    if this.S2.IsEmpty() {
        var tmp int
        for !this.S1.IsEmpty() {
            tmp = this.S1.Pop()
            this.S2.Push(tmp)
        }
    }
    return this.S2.Pop()
}


/** Get the front element. */
func (this *MyQueue) Peek() int {
    if this.S2.IsEmpty() {
        var tmp int
        for !this.S1.IsEmpty() {
            tmp = this.S1.Pop()
            this.S2.Push(tmp)
        }
    }
    return this.S2.Top()
}


/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
    return this.S1.IsEmpty() && this.S2.IsEmpty()
}