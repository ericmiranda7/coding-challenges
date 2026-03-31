func shortestPathBinaryMatrix(grid [][]int) int {
    tgtNode := Node{len(grid)-1, len(grid)-1}
    q := []Node{}
    q = append(q, Node{0, 0})

    lvl := 1
    for len(q) > 0 {
        qlen := len(q)
        for _ = range qlen {
            node := q[0]
            q = q[1:]
            
            if node == tgtNode {
                return lvl
            }
            if grid[node[ROW]][node[COL]] == WALL {
                continue
            }


            nbrs := getNeighbours(grid, node)
            q = append(q, nbrs...)
            grid[node[ROW]][node[COL]] = 1 // mark visited
        }
        lvl += 1
    }

    return -1
}

func getNeighbours(grid [][]int, node Node) []Node {
    p := []Node{
        Node{node[ROW]-1, node[COL]-1}, Node{node[ROW]-1, node[COL]}, Node{node[ROW]-1, node[COL]+1},
        Node{node[ROW], node[COL]-1}, Node{node[ROW], node[COL]+1},
        Node{node[ROW]+1, node[COL]-1}, Node{node[ROW]+1, node[COL]}, Node{node[ROW]+1, node[COL]+1},
    }

    res := []Node{}
    for _, n := range p {
        if n[ROW] < 0 || n[ROW] >= len(grid) {
            continue
        }
        if n[COL] < 0 || n[COL] >= len(grid) {
            continue
        }
        if grid[n[ROW]][n[COL]] == WALL {
            continue
        }

        res = append(res, n)
    }

    return res
}

type Node [2]int
const ROW = 0
const COL = 1
const VISITABLE = 0
const WALL = 1

// 1 0 0
// 1 1 0
// 1 1 0
