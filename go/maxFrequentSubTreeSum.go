/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
var counts map[int]int
var max int

func findFrequentTreeSum(root *TreeNode) []int {
    counts = make(map[int]int)
    max = 0
    result := make([]int, 0)
    findSum(root)
    for k, v := range counts {
        if v == max {
            result = append(result, k)
        }
    }
    return result
}

func findSum(root *TreeNode) int {
    if root == nil {
        return 0
    }
    sum := root.Val + findSum(root.Left) + findSum(root.Right)
    counts[sum]++
    if counts[sum] > max {
        max = counts[sum]
    }
    return sum
}