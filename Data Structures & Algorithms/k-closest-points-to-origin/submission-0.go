func kClosest(points [][]int, k int) [][]int {
    // make pq of size k
    // for each point
        // insert into pq
    // return pq
    ih := &ItemHeap{}
    heap.Init(ih)

    for _, point := range points {
        item := Item{point, calcDistance(point)}
        heap.Push(ih, item)

        if ih.Len() > k {
            heap.Pop(ih)
        }
    }

    res := [][]int{}
    for _, item := range *ih {
        res = append(res, item.point)
    }
    return res
}

type ItemHeap []Item

type Item struct {
    point []int
    distance float64
}

func (ih ItemHeap) Len() int {
    return len(ih)
}

func (ih ItemHeap) Less(i, j int) bool {
    return ih[i].distance > ih[j].distance
}

func (ih ItemHeap) Swap(i, j int) {
    ih[i], ih[j] = ih[j], ih[i]
}

func (ih *ItemHeap) Push(x any) {
    item := x.(Item)
    *ih = append(*ih, item)
}

func (ih *ItemHeap) Pop() any {
    old := *ih
    res := old[len(old)-1]
    *ih = old[:len(old)-1]
    return res
}

func calcDistance(point []int) float64 {
    in := math.Pow(0.0 - float64(point[0]), 2) + math.Pow(0.0 - float64(point[1]), 2)
    return math.Sqrt(in)
}