/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    return helper(nums, 0, len(nums) - 1)
}

func helper(nums []int, start, end int) *TreeNode {
    if start > end {
        return nil
    }
    mid := start + (end - start) / 2
    root := &TreeNode{
        Val: nums[mid],
    }
    root.Left = helper(nums, start, mid - 1)
    root.Right = helper(nums, mid + 1, end)
    return root
}