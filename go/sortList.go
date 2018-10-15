func sortList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    newHead := head
    head = head.Next
    newHead.Next = nil 
    for head != nil {
        tmp := head
        head = head.Next
        if tmp.Val < newHead.Val {
            tmp.Next = newHead
            newHead = tmp
        } else {
            p, c := newHead, newHead
            for c != nil && c.Val < tmp.Val {
                p = c
                c = c.Next
            }
            tmp.Next = p.Next
            p.Next = tmp
        }
    }
    return newHead
}