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
        // delete
        // if 0..1 child
        if root.Left == nil {
            return root.Right
        }
        if root.Right == nil {
            return root.Left
        }

        // if 2 children
        minVal := findMinIn(root.Right)
        root.Val = minVal
        root.Right = deleteNode(root.Right, minVal)
        return root
    } else if key < root.Val {
        // move left
        root.Left = deleteNode(root.Left, key)
    } else {
        // move right
        root.Right = deleteNode(root.Right, key)
    }

    return root
}

func findMinIn(root *TreeNode) int {
    if root.Left == nil {
        return root.Val
    }

    return findMinIn(root.Left)
}
