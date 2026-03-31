type DynamicArray struct {
    backingArr []int
    size int
    cap int
}

func NewDynamicArray(capacity int) *DynamicArray {
    arr := make([]int, capacity)
    return &DynamicArray{arr, 0, capacity}
}

func (da *DynamicArray) Get(i int) int {
    return da.backingArr[i]
}

func (da *DynamicArray) Set(i int, n int) {
    da.backingArr[i] = n
}

func (da *DynamicArray) Pushback(n int) {
    if da.size == da.cap {
        da.resize()
    }

    da.backingArr[da.size] = n
    da.size += 1
}

func (da *DynamicArray) Popback() int {
    res := da.backingArr[da.size - 1]
    da.backingArr = da.backingArr[0:da.size - 1]
    da.size -= 1
    return res
}

func (da *DynamicArray) resize() {
    newArr := make([]int, da.cap * 2)
    for i, _ := range da.backingArr {
        newArr[i] = da.backingArr[i]
    }
    da.backingArr = newArr
    da.cap *= 2
}

func (da *DynamicArray) GetSize() int {
    return da.size
}

func (da *DynamicArray) GetCapacity() int {
    return da.cap
}
