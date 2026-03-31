/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
    var prev *ListNode
    
    for itr := head; itr != nil; {
        tmp := itr.Next
        itr.Next = prev
        prev = itr
        itr = tmp
    }
    return prev
}
