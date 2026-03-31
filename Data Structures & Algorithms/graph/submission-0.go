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
    visited := map[int]bool{}
    return dfs(src, dst, g, visited)
}

func dfs(src, dst int, g *Graph, visited map[int]bool) bool {
    if src == dst {
        return true
    }
    if visited[src] {
        return false
    }

    visited[src] = true
    for _, dest := range g.adjList[src] {
        if dfs(dest, dst, g, visited) {
            return true
        }
    }
    return false
}
