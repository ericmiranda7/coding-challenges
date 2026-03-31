/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func kthSmallest(root *TreeNode, k int) int {
    // ksmallest(root, k, count, res)   
        // if root == nil
            // return count

        // count = ksmallest(root.Left, k, count, res) + 1
        // if count == k
            // res.add(root.Val)
        // return ksmallest(root.Right, k, count, res)
    
    res := []int{}
    ksmallest(root, k, 0, &res)
    return res[0]
}

func ksmallest(root *TreeNode, k, count int, res *[]int) int {
    if root == nil {
        return count
    }

    count = ksmallest(root.Left, k, count, res) + 1
    if count == k {
        fmt.Println("yaay", root.Val)
        *res = append(*res, root.Val)
        return count + 1
    }
    return ksmallest(root.Right, k, count, res)
}
