type HashTable struct {
    arr []*Item
    capacity int
    size int
}

type Item struct {
    key int
    value int
}

func NewHashTable(capacity int) *HashTable {
    arr := make([]*Item, capacity)
    return &HashTable{arr, capacity, 0}
}

func (ht *HashTable) hash(key int) int {
    return key % ht.capacity
}

func (ht *HashTable) Insert(key, value int) {
    item := &Item{key, value}

    indx := ht.getIndex(key)
    if indx == -1 {
        // insert
        hash := ht.hash(key)

        for ht.arr[hash] != nil {
            hash = (hash + 1) % ht.capacity
        }

        ht.arr[hash] = item
        ht.size += 1

        if float32(ht.size) / float32(ht.capacity) >= 0.5 {
            ht.Resize()
        }
    } else {
        ht.arr[indx].value = value
    }
}

func (ht *HashTable) Get(key int) int {
    indx := ht.getIndex(key)
    
    if indx == -1 {
        return -1
    }

    return ht.arr[indx].value
}

func (ht *HashTable) getIndex(key int) int {
    hash := ht.hash(key)
    if ht.arr[hash] == nil {
        return -1
    }

    for hash < ht.capacity && ht.arr[hash] != nil {
        if ht.arr[hash].key == key {
            return hash
        }
        hash++
    }

    return -1
}

func (ht *HashTable) Remove(key int) bool {
    indx := ht.getIndex(key)
    if indx == -1 {
        return false
    }

    ht.arr[indx] = nil
    ht.size -= 1
    return true
}

func (ht *HashTable) GetSize() int {
    return ht.size
}

func (ht *HashTable) GetCapacity() int {
    return ht.capacity
}

func (ht *HashTable) Resize() {
    ht.capacity *= 2
    res := make([]*Item, ht.capacity)

    old := ht.arr
    ht.arr = res
    ht.size = 0
    for _, item := range old {
        if item == nil {
            continue
        }
        ht.Insert(item.key, item.value)
    }
}
