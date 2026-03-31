/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func insertIntoBST(root *TreeNode, val int) *TreeNode {
    itr := root
    if itr == nil {
        return &TreeNode{val, nil, nil}
    }
    for ;; {
        if val < itr.Val {
            fmt.Println("going left", itr.Val)
            if itr.Left == nil {
                itr.Left = &TreeNode{val, nil, nil}
                break
            }
            itr = itr.Left
        } else {
            fmt.Println("going left", itr.Val)
            if itr.Right == nil {
                itr.Right = &TreeNode{val, nil, nil}
                break
            }
            itr = itr.Right
        }
    }

    return root
}
