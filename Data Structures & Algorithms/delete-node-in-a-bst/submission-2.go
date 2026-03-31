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
        if root.Left == nil || root.Right == nil {
            // 0..1 child
            if root.Left == nil {
                return root.Right
            }
            return root.Left
        } else {
            // 2 chidlren
            minNode := findMin(root.Right)
            fmt.Println("min is", minNode)
            deleteNode(root, minNode.Val)
            root.Val = minNode.Val
            return root
        }
    }

    // rec case
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else {
        root.Right = deleteNode(root.Right, key)
    }

    return root
}

func findMin(root *TreeNode) *TreeNode {
    if root.Left == nil {
        return root
    }

    return findMin(root.Left)
}