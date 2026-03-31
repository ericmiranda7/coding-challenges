type LinkedList struct {
    head *Node
    tail *Node
    size int
}

type Node struct {
    next *Node
    val int
}

func NewLinkedList() *LinkedList {
    head := &Node{}
    tail := head
    return &LinkedList{head, tail, 0}
}

func (ll *LinkedList) Get(index int) int {
    if index >= ll.size {
        return -1
    }

    itr := ll.head.next
    for _ = range index {
        itr = itr.next
    }
    return itr.val
}

func (ll *LinkedList) InsertHead(val int) {
    nn := &Node{next: ll.head.next, val: val}

    if ll.tail == ll.head {
        // first ele
        ll.tail = nn
    }

    ll.head.next = nn
    ll.size += 1
}

func (ll *LinkedList) InsertTail(val int) {
    nn := &Node{next: nil, val: val}

    ll.tail.next = nn
    ll.tail = nn
    ll.size += 1
}

func (ll *LinkedList) Remove(index int) bool {
    if index < 0 || index >= ll.size {
        return false
    }

    itr := ll.head
    for _ = range index {
        itr = itr.next
    }

    // remove
    if itr.next == ll.tail {
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
