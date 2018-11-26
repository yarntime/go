/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
    value := make([]int, 0)
    for head != nil {
        value = append(value, head.Val)
        head = head.Next
    }
    for i := 0; i < len(value) / 2; i++ {
        if value[i] != value[len(value) - i - 1] {
            return false
        }
    }
    return true
}
