/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    res := []int{}

    addInOrder(&res, root)
    return res
}

func addInOrder(res *[]int, root *TreeNode) {
    if root == nil {
        return
    }

    addInOrder(res, root.Left)
    *res = append(*res, root.Val)
    addInOrder(res, root.Right)
}
