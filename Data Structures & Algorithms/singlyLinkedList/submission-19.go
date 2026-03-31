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
        &ListNode{},
        nil,
        0,
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
    newNode := &ListNode{val, nil}
    if (ll.head.next == nil) {
        // first ele
        ll.tail = newNode
    }
    newNode.next = ll.head.next
    ll.head.next = newNode

    ll.size += 1
}

func (ll *LinkedList) InsertTail(val int) {
    newNode := &ListNode{val, nil}
    if (ll.tail != nil) {
        ll.tail.next = newNode
    } else {
        ll.head.next = newNode
    }
    ll.tail = newNode

    ll.size += 1
}

func (ll *LinkedList) Remove(index int) bool {
    if (index >= ll.size) {
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
    res := []int{}
    
    itr := ll.head.next
    for itr != nil {
        res = append(res, itr.val)
        itr = itr.next
    }

    return res
}
