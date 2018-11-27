/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 || k == 1 {
		return head
	}
	result := &ListNode{}
	p := result
	h := head
	e := head
	for {
		n := k - 1
		for n != 0 && e != nil {
			n --
			e = e.Next
		}
		if n != 0 || e == nil {
			p.Next = h
			break
		}
		newh := e.Next
		e.Next = nil
		ns, ne := reverseList(h)
		p.Next = ns
		p = ne
		h = newh
		e = newh
	}
	return result.Next
}

func reverseList(head *ListNode) (*ListNode, *ListNode) {
	start := head
	end := head
	head = head.Next
	start.Next = nil
	for head != nil {
		t := head.Next
		head.Next = start
		start = head
		head = t
	}
	return start, end
}