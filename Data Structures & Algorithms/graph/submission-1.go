import "slices"

type Graph struct {
    adjList map[int][]int
}

func NewGraph() *Graph {
    return &Graph{map[int][]int{}}
}

func (g *Graph) AddEdge(src, dst int) {
    if _, ok := g.adjList[src]; !ok {
        g.adjList[src] = []int{}
    }
    if _, ok := g.adjList[dst]; !ok {
        g.adjList[dst] = []int{}
    }

    for _, dest := range g.adjList[src] {
        if dst == dest {
            return
        }
    }
    g.adjList[src] = append(g.adjList[src], dst)
}

func (g *Graph) RemoveEdge(src, dst int) bool {
    if _, ok := g.adjList[src]; !ok {
        return false
    }

    for i, dest := range g.adjList[src] {
        if dest == dst {
            g.adjList[src] = slices.Concat(g.adjList[src][0:i], g.adjList[src][i+1:])
            return true
        }
    }

    return false
}

func (g *Graph) HasPath(src, dst int) bool {
    return bfs(src, dst, g)
}

func bfs(src, dst int, g *Graph) bool {
    q := []int{}
    q = append(q, src)
    visited := map[int]bool{}

    for len(q) != 0 {
        node := q[0]
        q = q[1:]
        visited[node] = true
        
        for _, nbr := range g.adjList[node] {
            if nbr == dst {
                return true
            }
            if !visited[nbr] {
                q = append(q, nbr)
            }
        }
    }

    return false
}
