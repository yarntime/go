type Queue struct {
    element []*TreeNode
}

func NewQueue() *Queue {
    return &Queue{
        element: make([]*TreeNode, 0),
    }
}

func (q *Queue) put(node *TreeNode) {
    q.element = append(q.element, node)
}

func (q *Queue) take() *TreeNode {
    node := q.element[0]
    q.element = q.element[1:]
    return node
}

func (q *Queue) isEmpty() bool {
    return len(q.element) == 0
}

func averageOfLevels(root *TreeNode) []float64 {
    result := []float64{}
    q := NewQueue()
    q.put(root)
    last := root
    curLast := root
    curCount := 0
    curVal := 0
    for !q.isEmpty() {
        n := q.take()
        if n.Left != nil {
            q.put(n.Left)
            curLast = n.Left
        }
        if n.Right != nil {
            q.put(n.Right)
            curLast = n.Right
        }
        curVal += n.Val
        curCount += 1
        if n == last {
            result = append(result, float64(curVal) / float64(curCount))
            curCount = 0
            curVal = 0
            last = curLast
        }
    }
    return result
}