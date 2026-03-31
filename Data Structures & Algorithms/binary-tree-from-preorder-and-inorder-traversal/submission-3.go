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
    nn := &TreeNode{preorder[0], nil, nil}
    if len(preorder) == 1 {
        return nn
    }

    partitionIndx := getPartitionIndx(preorder[0], inorder)

    fmt.Println("hi")
    leftPreorder := preorder[1:partitionIndx+1]
    rightPreorder := preorder[partitionIndx+1:]
    fmt.Println("bye")

    fmt.Println("ok")
    nn.Left = buildTree(leftPreorder, inorder[0:partitionIndx])
    nn.Right = buildTree(rightPreorder, inorder[partitionIndx+1:])
    fmt.Println("eok")

    return nn
}

func getPartitionIndx(needle int, haystack []int) int {
    fmt.Println("needle", needle, "hs", haystack)
    var indx int
    for i, v := range haystack {
        if needle == v {
            indx = i
            break
        }
    }

    return indx
}