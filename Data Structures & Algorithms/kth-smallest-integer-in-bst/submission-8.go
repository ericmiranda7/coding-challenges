/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var res int
    preorder(root, k, 0, &res)
    return res
}

func preorder(root *TreeNode, k, count int, res *int) int {
    if root == nil {
        return count
    }

    count = preorder(root.Left, k, count, res) + 1
    if count == k {
        *res = root.Val
    }
    return preorder(root.Right, k, count, res)
}
