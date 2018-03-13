func swapPairs(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    n := head.Next
    head.Next = swapPairs(head.Next.Next)
    n.Next = head
    return n
}