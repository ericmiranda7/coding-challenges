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
	if len(lists) == 1 {
		return lists[0]
	}

	m := len(lists) / 2
	lhalf := mergeKLists(lists[0:m])
	rhalf := mergeKLists(lists[m:])
	return merge(lhalf, rhalf)
}

func merge(lhalf, rhalf *ListNode) *ListNode {
	head := &ListNode{}
	itr := head

	for lhalf != nil && rhalf != nil {
		if lhalf.Val < rhalf.Val {
			itr.Next = lhalf
			lhalf = lhalf.Next
		} else {
			itr.Next = rhalf
			rhalf = rhalf.Next
		}
		itr = itr.Next
	}

	if lhalf == nil {
		itr.Next = rhalf
	} else {
		itr.Next = lhalf
	}

	return head.Next
}
