func findKthLargest(nums []int, k int) int {
    // mh
    // for each num
        // insert into mh
        // if len(mh) > k
            // mh.pop
    
    // return mh.pop
    mh := &IntHeap{}
    heap.Init(mh)

    for _, num := range nums {
        heap.Push(mh, num)
        if mh.Len() > k {
            heap.Pop(mh)
        }
    }

    return heap.Pop(mh).(int)
}

type IntHeap []int

func (ih IntHeap) Len() int {
    return len(ih)
}

func (ih IntHeap) Less(i, j int) bool {
    return ih[i] < ih[j]
}

func (ih IntHeap) Swap(i, j int) {
    ih[i], ih[j] = ih[j], ih[i]
}

func (ih *IntHeap) Push(x any) {
    *ih = append(*ih, x.(int))
}

func (ih *IntHeap) Pop() any {
    old := *ih
    x := old[len(old)-1]
    *ih = old[0:len(old)-1]
    return x
}
