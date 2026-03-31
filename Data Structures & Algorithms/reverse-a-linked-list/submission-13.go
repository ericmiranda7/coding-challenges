/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    if head == nil {
        return head
    }
    cur := head
    var prev *ListNode
    for cur != nil {
        tmp := cur.Next
        cur.Next = prev
        prev = cur
        cur = tmp
    }

    return prev
}
