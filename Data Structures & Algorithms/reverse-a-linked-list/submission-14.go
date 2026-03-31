/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// 0 - 1 - 2 - 3 -
// 3 - 2 - 1 - 0

func reverseList(head *ListNode) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }

    newHead := reverseList(head.Next)
    head.Next.Next = head
    head.Next = nil

    return newHead
}
