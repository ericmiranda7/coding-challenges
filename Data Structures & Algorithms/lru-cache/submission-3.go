type LRUCache struct {
    head *Node
    tail *Node
    capacity int
    size int
    hmap map[int]*Node
}

type Node struct {
    prev *Node
    next *Node
    value int
    key int
}

func Constructor(capacity int) LRUCache {
    head := &Node{nil, nil, -1, -1}
    tail := &Node{nil, nil, -1, -1}
    head.next = tail
    tail.prev = head
    return LRUCache{head, tail, capacity, 0, map[int]*Node{}}
}

func (this *LRUCache) Get(key int) int {
    entry, ok := this.hmap[key]
    if !ok {
        return -1
    }

    this.update(entry)
    return entry.value
}

func (this *LRUCache) Put(key int, value int) {
    entry, exists := this.hmap[key]
    if exists {
        entry.value = value
        this.update(entry)
    } else {
        if this.size >= this.capacity {
            entry := this.evict()
            delete(this.hmap, entry.key)
        }
        this.add(key, value)
    }
}

func (this *LRUCache) update(entry *Node) {
    entry.prev.next = entry.next
    entry.next.prev = entry.prev

    this.head.next.prev = entry
    entry.next = this.head.next
    entry.prev = this.head
    this.head.next = entry
}

func (this *LRUCache) evict() *Node {
    entry := this.tail.prev
    entry.prev.next = this.tail
    this.tail.prev = entry.prev
    this.size -= 1
    return entry
}

func (this *LRUCache) add(key, value int) {
    n := &Node{this.head, this.head.next, value, key}
    this.head.next.prev = n
    this.head.next = n
    this.hmap[key] = n
    this.size += 1
}