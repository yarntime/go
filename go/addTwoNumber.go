/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
    c := 0
    result := &ListNode{}
    head := result
    for l1 != nil && l2 != nil {
        t := l1.Val + l2.Val + c
        l1 = l1.Next
        l2 = l2.Next
        head.Next = &ListNode{Val: t % 10, Next: nil}
        head = head.Next
        c = t / 10
    }
    for l1 != nil {
        t := l1.Val + c
        l1 = l1.Next
        head.Next = &ListNode{Val: t % 10, Next: nil}
        head = head.Next
        c = t / 10
    }
    for l2 != nil {
        t := l2.Val + c
        l2 = l2.Next
        head.Next = &ListNode{Val: t % 10, Next: nil}
        head = head.Next
        c = t / 10
    }
    if c != 0 {
        head.Next = &ListNode{Val: c, Next: nil}
    }
    return result.Next
    
}