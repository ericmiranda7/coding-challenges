/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    res := [][]int{}
    dfs(root, &res, 0)
    return res
}

func dfs(root *TreeNode, res *[][]int, depth int) {
    if root == nil {
        return
    }
    if len(*res) == depth {
        *res = append(*res, []int{})
        fmt.Println(res)
    }

    (*res)[depth] = append((*res)[depth], root.Val)
    dfs(root.Left, res, depth+1)
    dfs(root.Right, res, depth+1)
}
