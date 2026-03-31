/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func buildTree(preorder []int, inorder []int) *TreeNode {
    if len(preorder) == 0 {
        return nil
    }

    root := preorder[0]
    // inorder indx
    inorderIndex := 0
    for i, v := range inorder {
        if v == root {
            inorderIndex = i
        }
    }

    n := &TreeNode{Val: root}
    n.Left = buildTree(preorder[1:1+inorderIndex], inorder[0:inorderIndex])
    n.Right = buildTree(preorder[inorderIndex+1:], inorder[1+inorderIndex:])
    return n
}
