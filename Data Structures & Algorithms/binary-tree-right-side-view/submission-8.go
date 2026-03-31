/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rightSideView(root *TreeNode) []int {
    res := []int{}

    dfs(root, &res, 0)
    return res
}

func dfs(root *TreeNode, res *[]int, depth int) {
    if root == nil {
        return
    }
    if len(*res) == depth {
        *res = append(*res, root.Val)
    }

    dfs(root.Right, res, depth+1)
    dfs(root.Left, res, depth+1)
}
