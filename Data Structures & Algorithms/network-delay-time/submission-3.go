func networkDelayTime(times [][]int, n int, k int) int {
	adjMap := map[int][]*Edge{}
    for _, edge := range times {
		u, v, w := edge[0], edge[1], edge[2]
		adjMap[u] = append(adjMap[u], &Edge{u, v, w})
	}
	latencyMap := map[int]int{}
	for i := 1; i <= n; i++ {
		latencyMap[i] = math.MaxInt
	}

	pq := &PriorityQueue{}
	heap.Push(pq, &Item{&Node{k, adjMap[k]}, 0, 0})

	for len(*pq) > 0 {
		item := heap.Pop(pq).(*Item)

		if item.priority >= latencyMap[item.value.id] {
			continue
		}

		latencyMap[item.value.id] = item.priority

		for _, edge := range item.value.edges {
			if item.priority + edge.cost < latencyMap[edge.dest] {
				heap.Push(pq, &Item{&Node{edge.dest, adjMap[edge.dest]}, item.priority + edge.cost, 0})
			}
		}
	}

	res := -1
	for _, v := range latencyMap {
		if v == math.MaxInt {
			return -1
		} else if v > res {
			res = v
		}
	}

	return res
}

type Node struct {
	id int
	edges []*Edge
}

type Edge struct {
	source int
	dest int
	cost int
}

// An Item is something we manage in a priority queue.
type Item struct {
	value    *Node // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // don't stop the GC from reclaiming the item eventually
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}