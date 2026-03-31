/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    inOrderArr := inOrder(root)
    return inOrderArr[k-1]
}

func inOrder(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    res := inOrder(root.Left)
    res = append(res, root.Val)
    res = append(res, inOrder(root.Right)...)
    return res
}
