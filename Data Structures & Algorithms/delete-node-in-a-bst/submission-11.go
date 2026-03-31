/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
 
func deleteNode(root *TreeNode, key int) *TreeNode {
    deleteNode := root
    var parentNode *TreeNode

    for deleteNode != nil {
        
        if deleteNode.Val == key {
            break
        }

        parentNode = deleteNode
        
        if key < deleteNode.Val {
            deleteNode = deleteNode.Left
        } else {
            deleteNode = deleteNode.Right
        }
    }

    if deleteNode == nil {
        return root
    } 

    // 0..1 child
    if deleteNode.Left == nil {
        if parentNode != nil {
            if parentNode.Left == deleteNode {
                parentNode.Left = deleteNode.Right
            } else {
                parentNode.Right = deleteNode.Right
            }
        } else {
            return deleteNode.Right
        }
        return root
    }
    if deleteNode.Right == nil {
        if parentNode != nil {
            if parentNode.Left == deleteNode {
                parentNode.Left = deleteNode.Left
            } else {
                parentNode.Right = deleteNode.Left
            }
        } else {
            return deleteNode.Left
        }
        return root
    }

    fmt.Println(deleteNode.Val)

    // 0..2 children
    minNode := deleteNode.Right
    minParent := deleteNode

    for minNode.Left != nil {
        minParent = minNode
        minNode = minNode.Left
    }
    deleteNode.Val = minNode.Val
    if minParent == deleteNode {
        // didnt move
        deleteNode.Right = minNode.Right
    } else {
        minParent.Left = minNode.Right
    }

    return root

}
