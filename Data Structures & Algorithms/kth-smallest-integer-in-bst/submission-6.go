/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    var ans int = -1
    inOrder(root, k, 0, &ans)
    return ans
}

func inOrder(root *TreeNode, k, ws int, ans *int) int {
    if root == nil {
        return ws
    }

    ws = inOrder(root.Left, k, ws, ans) + 1
    if ws == k {
        // mark answer
        fmt.Println("setting ", ws, k, root.Val)
        *ans = root.Val
    } else if *ans == -1 {
        ws = inOrder(root.Right, k, ws, ans)
    }

    return ws
}