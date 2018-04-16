/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func isBalanced(root *TreeNode) bool {
    if root == nil {
        return true
    }
    _, _, result := helper(root)
    return result
}

func helper(root *TreeNode) (int, int, bool) {
    if root == nil {
        return 0, 0, true
    }
    leftMin, leftMax, leftResult := helper(root.Left)
    rightMin, rightMax, rightResult := helper(root.Right)
    if leftResult && rightResult {
        leftMin, leftMax, rightMin, rightMax = leftMin + 1, leftMax + 1, rightMin + 1, rightMax + 1
        diff1 := abs(leftMin, rightMax)
        diff2 := abs(leftMax, rightMin)
        if diff1 <= 1 && diff2 <= 1 {
            return min(leftMin, rightMin), max(leftMax, rightMax), true
        } else {
            return -1, -1, false
        }
    } else {
        return -1, -1, false
    }
}

func max(a, b int) int {
    if a > b {
        return a
    }
    return b
}

func min(a, b int) int {
    if a < b {
        return a
    }
    return b
}

func abs(a, b int) int {
    c := a - b
    if c < 0 {
        return -c
    }
    return c
}