type LinkedList struct {
    head *ListNode
    tail *ListNode
    size int
}

type ListNode struct {
    next *ListNode
    val int
}

func NewLinkedList() *LinkedList {
    baseNode := &ListNode{nil, -1}
    return &LinkedList{baseNode, baseNode, 0}
}

func (ll *LinkedList) Get(index int) int {
    if index >= ll.size {
        return -1
    }

    itr := ll.head.next
    for i := 0; i < index; i++ {
        itr = itr.next
    }

    return itr.val
}

func (ll *LinkedList) InsertHead(val int) {
    nn := &ListNode{ll.head.next, val}
    if (ll.size == 0) {
        // first ele
        ll.tail = nn
    }
    ll.head.next = nn

    ll.size += 1
}

func (ll *LinkedList) InsertTail(val int) {
    nn := &ListNode{nil, val}
    if (ll.size == 0) {
        ll.head.next = nn
    } else {
        ll.tail.next = nn
    }
    ll.tail = nn

    ll.size += 1
}

func (ll *LinkedList) Remove(index int) bool {
    if index >= ll.size {
        return false
    }

    itr := ll.head
    for i := 0; i < index; i++ {
        itr = itr.next
    }
    if (itr.next == ll.tail) {
        ll.tail = itr
    }

    itr.next = itr.next.next
    ll.size -= 1
    return true
}

func (ll *LinkedList) GetValues() []int {
    itr := ll.head.next
    res := []int{}
    for itr != nil {
        res = append(res, itr.val)
        itr = itr.next
    }

    return res
}
