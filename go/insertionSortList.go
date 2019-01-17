/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
    if head == nil {
        return nil
    }
    result := head
    head = head.Next
    result.Next = nil
    
    for head != nil {
        t := head
        head = head.Next
        if t.Val <= result.Val {
            t.Next = result
            result = t
        } else {
            cur := result
            for cur.Next != nil {
                if cur.Next.Val >= t.Val {
                    break
                }
                cur = cur.Next
            }
            t.Next = cur.Next
            cur.Next = t
        }
    }
    return result
}