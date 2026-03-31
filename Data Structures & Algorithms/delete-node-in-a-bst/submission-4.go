/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return root
    }
    if root.Val == key {
        // delete this
        if root.Right == nil || root.Left == nil {
            if root.Right == nil {
                return root.Left
            }
            return root.Right
        }

        max := getMaxFrom(root.Left)
        deleteNode(root, max.Val)
        root.Val = max.Val
        return root
    } 

    if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        root.Left = deleteNode(root.Left, key)
    }
    return root
}

func getMaxFrom(root *TreeNode) *TreeNode {
    if root.Right == nil {
        return root
    }
    return getMaxFrom(root.Right)
}