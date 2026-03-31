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

func mergeHelper(lists []*ListNode, s int, e int) *ListNode {
    if e <= s {
        return lists[s]
    }

    m := ((e - s) / 2) + s
    llist := mergeHelper(lists, s, m)
    rlist := mergeHelper(lists, m + 1, e)
    return merge(llist, rlist)
}

func merge(llist, rlist *ListNode) *ListNode {
    res := &ListNode{-1, nil}
    itr := res

    for llist != nil && rlist != nil {
        if llist.Val <= rlist.Val {
            itr.Next = llist
            llist = llist.Next
        } else {
            itr.Next = rlist
            rlist = rlist.Next
        }
        itr = itr.Next
    }

    for llist != nil {
        itr.Next = llist
        llist = llist.Next
        itr = itr.Next
    }
    for rlist != nil {
        itr.Next = rlist
        rlist = rlist.Next
        itr = itr.Next
    }

    return res.Next
}
