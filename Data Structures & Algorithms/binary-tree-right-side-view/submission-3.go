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

    for len(q) != 0 {
        lvl := []int{}
        lvln := len(q)

// wb ;
        for ;lvln > 0; lvln-- {
            node := q[0]
            q = q[1:]

            lvl = append(lvl, node.Val)
            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }
        }
        res = append(res, lvl[len(lvl)-1])
    }

    return res
}
