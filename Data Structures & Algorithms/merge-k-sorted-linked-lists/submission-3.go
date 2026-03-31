/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// todo(): rec sol?
func mergeKLists(lists []*ListNode) *ListNode {
    // if lists.empty return lists
    // ptr goes from [1]..end
        // merge([0], lists[ptr])
    // return lists[0]
    
    if len(lists) == 0 {
        return nil
    }

    for ptr := 1; ptr < len(lists); ptr++ {
        lists[0] = merge(lists[0], lists[ptr])
    }
    return lists[0]
}

func merge(res *ListNode, val *ListNode) *ListNode {
    dummy := &ListNode{}
    final := dummy

    var lptr = res
    var rptr = val
    for lptr != nil && rptr != nil {
        if rptr.Val < lptr.Val {
            dummy.Next = rptr
            rptr = rptr.Next
        } else {
            dummy.Next = lptr
            lptr = lptr.Next
        }
        dummy = dummy.Next
    }
    if lptr == nil {
        dummy.Next = rptr
    } else {
        dummy.Next = lptr
    }
    return final.Next
}