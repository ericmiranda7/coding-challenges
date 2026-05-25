func shortestPath(n int, edges [][]int, src int) map[int]int {
    resMap := map[int]int{}
    adjMap := map[int][]*Edge{}
    for _, edge := range edges {
        u, v, w := edge[0], edge[1], edge[2]
        adjMap[u] = append(adjMap[u], &Edge{u, v, w})
    }
    for i := 0; i < n; i++ {
        resMap[i] = math.MaxInt
    }
	pq := PriorityQueue{}
    heap.Push(&pq, &Item{&Node{adjMap[src], src}, 0, 0})

    for len(pq) > 0 {
        item := heap.Pop(&pq).(*Item)

        documentedCost := resMap[item.node.id]
        if item.priority < documentedCost {
            resMap[item.node.id] = item.priority
        }

        for _, edge := range item.node.edges {
            if item.priority + edge.weight < resMap[edge.dest] {
                heap.Push(&pq, &Item{&Node{adjMap[edge.dest], edge.dest}, item.priority + edge.weight, 0})
            }
        }
    }

    for k, v := range resMap {
        if v == math.MaxInt {
            resMap[k] = -1
        }
    }

    return resMap
}

type Edge struct {
    source int
    dest int
    weight int
}

type Node struct {
    edges []*Edge
    id int
}

// An Item is something we manage in a priority queue.
type Item struct {
	node *Node // The value of the item; arbitrary.
	priority int    // The priority of the item in the queue.
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {

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