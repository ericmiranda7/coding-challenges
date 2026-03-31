/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func levelOrder(root *TreeNode) [][]int {
    if root == nil {
        return [][]int{}
    }
    res := [][]int{}
    q := list.New()
    q.PushBack(root)

    for q.Len() != 0 {
        subRes := []int{}
        nodes := []*TreeNode{}

        for q.Len() != 0 {
            n, ok := q.Remove(q.Front()).(*TreeNode)
            if !ok {
                fmt.Println("cant cast")
            }
            nodes = append(nodes, n)
        }

        for _, node := range nodes {
            subRes = append(subRes, node.Val)

            if node.Left != nil {
                q.PushBack(node.Left)
            }
            if node.Right != nil {
                q.PushBack(node.Right)
            }
        }
        res = append(res, subRes)
    }

    return res
}
