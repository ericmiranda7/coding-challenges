type Deque struct {
    head *Node
    tail *Node
}

type Node struct {
    val int
    prev *Node
    next *Node
}


func NewDeque() *Deque {
    head := &Node{-1, nil, nil}
    tail := &Node{-1, nil, nil}
    head.next = tail
    tail.prev = head

    return &Deque{head, tail}
}

func (d *Deque) IsEmpty() bool {
    return d.head.next == d.tail
}

func (d *Deque) Append(value int) {
    nn := &Node{value, d.tail.prev, d.tail}
    d.tail.prev.next = nn
    d.tail.prev = nn
}

func (d *Deque) AppendLeft(value int) {
    nn := &Node{value, d.head, d.head.next}
    d.head.next.prev = nn
    d.head.next = nn
}

func (d *Deque) Pop() int {
    if d.head.next == d.tail {
        return -1
    }

    x := d.tail.prev
    d.tail.prev.prev.next = d.tail
    d.tail.prev = d.tail.prev.prev
    return x.val
}

func (d *Deque) PopLeft() int {
    if d.head.next == d.tail {
        return -1
    }

    x := d.head.next
    d.head.next.next.prev = d.head
    d.head.next = d.head.next.next
    return x.val
}
