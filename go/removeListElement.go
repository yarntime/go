func removeElements(head *ListNode, val int) *ListNode {
    for head != nil && head.Val == val {
        head = head.Next
    }
    if head == nil {
        return head
    }
	next := head
	for next != nil && next.Next != nil {
		if next.Next.Val == val {
			next.Next = next.Next.Next
        } else {
            next = next.Next
        }
	}
	return head
}