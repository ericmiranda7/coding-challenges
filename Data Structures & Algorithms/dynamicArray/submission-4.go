type DynamicArray struct {
    capacity int
    len int
    arr []int
}

func NewDynamicArray(capacity int) *DynamicArray {
    return &DynamicArray{capacity, 0, make([]int, capacity)}
}

func (da *DynamicArray) Get(i int) int {
    return da.arr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.arr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if (da.len == da.capacity) {
        da.resize()
    }
    da.arr[da.len] = n
    da.len += 1
}

func (da *DynamicArray) Popback() int {
    da.len -= 1
    return da.arr[da.len]
}

func (da *DynamicArray) resize() {
    da.capacity *= 2
    newArr := make([]int, da.capacity)
    copy(newArr, da.arr)
    da.arr = newArr
}

func (da *DynamicArray) GetSize() int {
    return da.len
}

func (da *DynamicArray) GetCapacity() int {
    return da.capacity
}
