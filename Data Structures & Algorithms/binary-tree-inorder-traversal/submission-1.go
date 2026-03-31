/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    return addInOrder(root)
}

func addInOrder(root *TreeNode) []int {
    if root == nil {
        return []int{}
    }

    res := addInOrder(root.Left)
    res = append(res, root.Val)
    res = append(res, addInOrder(root.Right)...)

    return res
}
