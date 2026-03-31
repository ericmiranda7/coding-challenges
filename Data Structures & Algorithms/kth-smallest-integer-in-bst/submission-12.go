/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    stk := []*TreeNode{}
    itr := root

    count := 0
    for itr != nil || len(stk) > 0 {
        for itr != nil {
            stk = append(stk, itr)
            itr = itr.Left
        }

        n := stk[len(stk)-1]
        stk = stk[:len(stk)-1]

        count++
        if count == k {
            return n.Val
        }
        if n.Right != nil {
            itr = n.Right
        }
    }

    return -1
}
