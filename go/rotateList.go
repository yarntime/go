/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
    if head == nil {
        return nil
    }
    l := 0
    end := head
    for end.Next != nil {
        l++
        end = end.Next
    }
    l++
    c := l - (k % l)
    if c == 0 {
        return head
    }
    end.Next = head
    p := head
    for c != 0 {
        p = head
        head = head.Next
        c--
    }
    p.Next = nil
    return head
}