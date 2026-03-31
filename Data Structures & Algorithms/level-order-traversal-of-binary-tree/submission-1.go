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
    if root == nil {
        return res
    }
    q := []*TreeNode{}
    q = append(q, root)

    for len(q) != 0 {
        innerres := []int{}
        qlen := len(q)

        for qlen > 0 {
            n := q[0]
            q = q[1:]
            innerres = append(innerres, n.Val)
            if n.Left != nil {
                q = append(q, n.Left)
            }
            if n.Right != nil {
                q = append(q, n.Right)
            }
            qlen -= 1
        }

        res = append(res, innerres)
    }

    return res
}
