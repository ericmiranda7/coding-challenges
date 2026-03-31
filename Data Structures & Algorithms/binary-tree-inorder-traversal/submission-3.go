/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func inorderTraversal(root *TreeNode) []int {
    // res
    // stk
    // stk.append(root)
    // itr = root

    // for itr != nil || stk.len > 0
        // while itr != nil
            // stk.append(itr)
            // itr = itr.Left

        // n = stk.pop()
        // res.append(n.Val)
        // if n.Right != nil
            // itr = n.Right

    res := []int{}
    stk := []*TreeNode{}
    itr := root

    for itr != nil || len(stk) > 0 {
        for itr != nil {
            stk = append(stk, itr)
            itr = itr.Left
        }

        n := stk[len(stk)-1]
        stk = stk[:len(stk)-1]
        res = append(res, n.Val)
        
        if n.Right != nil {
            itr = n.Right
        }
    }

    return res
}
