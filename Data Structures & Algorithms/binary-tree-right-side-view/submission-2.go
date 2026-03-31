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
    if root == nil {
        return res
    }
    q := []*TreeNode{}
    q = append(q, root)

    // [(1)]
    for len(q) > 0 {
        qLen := len(q)
        for i := 0; i < qLen; i++ {
            n := q[0]
            q = q[1:]
            if i == qLen - 1 {
                res = append(res, n.Val)
            }
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
        }
    }

    return res
}
