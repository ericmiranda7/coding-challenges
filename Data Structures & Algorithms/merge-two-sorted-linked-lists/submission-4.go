/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    res := &ListNode{-1, nil}
    itr := res
    for ;list1 != nil && list2 != nil; {
        if list1.Val <= list2.Val {
            itr.Next = list1
            list1 = list1.Next
        } else {
            itr.Next = list2
            list2 = list2.Next
        }
        itr = itr.Next
    }

    if list1 == nil {
        itr.Next = list2
    } else {
        itr.Next = list1
    }

    return res.Next
}
