type Deque struct {
    head *node
    tail *node
    size int
}

type node struct {
    prev *node
    next *node
    val int
}


func NewDeque() *Deque {
    head := &node{nil, nil, -1}
    return &Deque{head, head, 0}
}

func (d *Deque) IsEmpty() bool {
    return d.size == 0
}

func (d *Deque) Append(value int) {
    nn := &node{d.tail, nil, value}
    d.tail.next = nn
    d.tail = nn
    d.size += 1
}

func (d *Deque) AppendLeft(value int) {
    nn := &node{d.head, d.head.next, value}
    
    if d.head != d.tail {
        d.head.next.prev = nn
    } else {
        d.tail = nn
    }
    d.head.next = nn
    d.size += 1
}

func (d *Deque) Pop() int {
    if d.size == 0 {
        return -1
    }

    v := d.tail.val
    d.tail = d.tail.prev
    d.tail.next = nil
    d.size -= 1

    return v
}

func (d *Deque) PopLeft() int {
    if d.size == 0 {
        return -1
    }

    v := d.head.next.val
    if d.head.next == d.tail {
        d.tail = d.head
    }
    d.head.next = d.head.next.next

    d.size -= 1
    
    return v
}
