/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    // if root == nil; return

    // if val < root.Val
        // res insertIntoBST(root.Left, val)
        // if res == nil
            // res.Left = nn
    // else
        // res = insertIntoBST(root.Right, val)
        // if res == nil
            // res.Right = nn
    // return root

    if root == nil {
        return &TreeNode{val, nil, nil}
    }

    if val < root.Val {
        root.Left = insertIntoBST(root.Left, val)
    } else {
        root.Right = insertIntoBST(root.Right, val)
    }

    return root
}
