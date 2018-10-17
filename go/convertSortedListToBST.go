/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedListToBST(head *ListNode) *TreeNode {
    if head == nil {
        return nil
    }
    if head.Next == nil {
        return &TreeNode{Val: head.Val, Left: nil, Right: nil}
    }
    
    return buildTree(head, nil)
}

func buildTree(head *ListNode, end *ListNode) *TreeNode {
    if head == end {
        return nil
    }
    slow, fast := head, head
    for fast != end && fast.Next != end {
        slow = slow.Next
        fast = fast.Next.Next
    }
    root := &TreeNode{Val: slow.Val}
    root.Left = buildTree(head, slow)
    root.Right = buildTree(slow.Next, end)
    return root
}