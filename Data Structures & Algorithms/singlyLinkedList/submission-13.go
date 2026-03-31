type LinkedList struct {
    head *ListNode
    tail *ListNode
    size int
}

type ListNode struct {
    val int
    next *ListNode
}

func NewLinkedList() *LinkedList {
    return &LinkedList{
        head: &ListNode{},
        tail: nil,
        size: 0,
    }
}

func (ll *LinkedList) Get(index int) int {
    if (index >= ll.size) {
        return -1
    }

    itr := ll.head.next
    for i := 0; i < index; i++ {
        itr = itr.next
    }

    return itr.val
}

func (ll *LinkedList) InsertHead(val int) {
    ln := &ListNode{val: val}

    if (ll.head.next == nil) {
        ll.tail = ln
    }
    ln.next = ll.head.next
    ll.head.next = ln

    ll.size += 1
}

func (ll *LinkedList) InsertTail(val int) {
    ln := &ListNode{val: val}

    if (ll.head.next == nil) {
        // first ele
        ll.head.next = ln
    } else {
        ll.tail.next = ln
    }
    ll.tail = ln

    ll.size += 1
}

func (ll *LinkedList) Remove(index int) bool {
    if (index >= ll.size) {
        return false
    }

    // if (index == 0) {
    //     if (ll.tail == ll.head.next) {
    //         ll.tail = nil
    //     }
    //     ll.head.next = ll.head.next.next
    //     ll.size -= 1
    //     return true
    // }

    itr := ll.head
    for i := 0; i < index; i++ {
        itr = itr.next
    }
    if (ll.tail == itr.next) {
        ll.tail = itr
    }
    itr.next = itr.next.next
    ll.size -= 1
    return true
}

func (ll *LinkedList) GetValues() []int {
    res := []int{}

    for i := ll.head.next; i != nil; i = i.next {
        res = append(res, i.val)
    }
    return res
}
