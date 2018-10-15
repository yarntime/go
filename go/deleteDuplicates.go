func deleteDuplicates(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    newHead := &ListNode{Next: nil}
    isEqual := false
    newCur := newHead
    p, c := head, head.Next 
    for c != nil {
        if c.Val == p.Val {
            isEqual = true
        } else if isEqual {
            isEqual = false
        } else if !isEqual {
            p.Next = newCur.Next
            newCur.Next = p
            newCur = newCur.Next
        }
        p = c
        c = c.Next
    }
    if !isEqual {
        newCur.Next = p
        newCur.Next.Next = nil
    }
    return newHead.Next
}