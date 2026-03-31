type Deque struct {
    head *Node
    tail *Node
}

type Node struct {
    prev *Node
    next *Node
    val int
}


func NewDeque() *Deque {
    return &Deque{nil, nil}
}

func (d *Deque) IsEmpty() bool {
    if (d.head == nil) {
        return true
    }
    return false
}

func (d *Deque) Append(value int) {
    nn := &Node{d.tail, nil, value}
    if (d.head == nil) {
        // first ele
        d.head = nn
        d.tail = nn
    } else {
        d.tail.next = nn
        d.tail = d.tail.next
    }
}

func (d *Deque) AppendLeft(value int) {
    nn := &Node{nil, d.head, value}
    if (d.head == nil) {
        d.tail = nn
    } else {
        d.head.prev = nn
    }
    d.head = nn
}

func (d *Deque) Pop() int {
    if (d.head == nil) {
        return -1 
    }

    val := d.tail.val
    d.tail = d.tail.prev
    if (d.tail != nil) {
        d.tail.next = nil
    } else {
        d.head = nil
    }
    return val
}

func (d *Deque) PopLeft() int {
    if (d.head == nil) {
        return -1 
    }

    val := d.head.val
    d.head = d.head.next
    if (d.head != nil) {
        d.head.prev = nil
    } else {
        d.tail = nil
    }
    return val
}
