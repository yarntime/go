func removeNthFromEnd(head *ListNode, n int) *ListNode {
    _, node := helper(head, n)
    return node
}

func helper(head *ListNode, n int) (int, *ListNode) {
    if head == nil {
        return 0, nil
    }
    curNum, curNode := helper(head.Next, n)
    curNum = curNum + 1
    head.Next = curNode
    if curNum == n {
        return curNum, curNode
    }
    return curNum, head
}