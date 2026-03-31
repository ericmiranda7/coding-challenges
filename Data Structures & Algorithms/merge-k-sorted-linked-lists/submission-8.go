/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeKLists(lists []*ListNode) *ListNode {
    if len(lists) == 0 {
        return nil
    }
    return mergeHelper(lists, 0, len(lists) - 1)
}

func mergeHelper(lists []*ListNode, start, end int) *ListNode {
    if end == start {
        return lists[end]
    }

    mid := ((end - start) / 2) + start
    l := mergeHelper(lists, start, mid)
    r := mergeHelper(lists, mid+1, end)

    return merge(l, r)
}

func merge(l, r *ListNode) *ListNode {
    res := &ListNode{-1, nil}
    dummy := res

    for l != nil && r != nil {
        if l.Val <= r.Val {
            dummy.Next = l
            l = l.Next
        } else {
            dummy.Next = r
            r = r.Next
        }
        dummy = dummy.Next
    }

    dummy.Next = l
    if r != nil {
        dummy.Next = r
    }

    return res.Next
}
