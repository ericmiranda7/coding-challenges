func lastStoneWeight(stones []int) int {
    ih := &IntHeap{}
    heap.Init(ih)
    for _, stone := range stones {
        heap.Push(ih, stone)
    }

    for ih.Len() > 1 {
        s1 := heap.Pop(ih).(int)
        s2 := heap.Pop(ih).(int)

        fmt.Println(s1, s2)

        if s1 == s2 {
            continue
        } else {
            s1 = s1 - s2
        }
        heap.Push(ih, s1)
    }

    if ih.Len() >= 1 {
        return heap.Pop(ih).(int)
    } else {
        return 0
    }
}

type IntHeap []int

func (h IntHeap) Len() int {
    return len(h)
}

func (h IntHeap) Less(i, j int) bool {
    return h[i] > h[j]
}

func (h IntHeap) Swap(i, j int) {
    h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x any) {
    val := x.(int)
    *h = append(*h, val)
}

func (h *IntHeap) Pop() any {
    old := *h
    x := old[len(*h)-1]
    *h = old[0:len(*h)-1]
    return x
}
