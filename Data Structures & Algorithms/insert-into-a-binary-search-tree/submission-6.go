/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    curr := root

    if curr == nil {
        return &TreeNode{Val: val}
    }

    for curr != nil {
        if val < curr.Val {
            if curr.Left == nil {
                curr.Left = &TreeNode{Val: val}
                break
            }
            curr = curr.Left
        } else {
            if curr.Right == nil {
                curr.Right = &TreeNode{Val: val}
                break
            }
            curr = curr.Right
        }
    }

    return root
}
