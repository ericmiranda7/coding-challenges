const ROW = 0
const COL = 1
const LAND = 0
const WATER = 1
type Node [2]int

func shortestPath(grid [][]int) int {
    endGame := Node{len(grid)-1, len(grid[len(grid)-1])-1}
    q := []Node{}
    q = append(q, Node{0, 0})

    lvl := 0
    for len(q) > 0 {
        qlen := len(q)
        for _ = range qlen {
            node := q[0]
            q = q[1:]
            
            if grid[node[ROW]][node[COL]] == WATER {
                continue
            }

            if node == endGame {
                return lvl
            }

            // node is land
            neighbours := getNeighbours(grid, node)
            q = append(q, neighbours...)
            grid[node[ROW]][node[COL]] = WATER
        }
        lvl += 1
    }

    return -1
}

func getNeighbours(grid [][]int, node Node) []Node {
    possibilities := [4]Node{
        Node{node[ROW] - 1, node[COL]},
        Node{node[ROW], node[COL] + 1}, Node{node[ROW], node[COL] - 1},
        Node{node[ROW] + 1, node[COL]},
    }
    res := []Node{}
    for _, n := range possibilities {
        if n[ROW] < 0 || n[ROW] >= len(grid) {
            continue
        }
        if n[COL] < 0 || n[COL] >= len(grid[n[ROW]]) {
            continue
        }
        res = append(res, n)

    }
    return res
}

// q = Node{0,0}
// lvl := 0
// for q != empty
    // qlen = len(q)
    // for i := range qlen
        // node := q.deq()
        // if node == WATER
            // continue
        // if node == ENDGAME
            // return lvl
        // node is land
        // neighbours = getNeighbours(node)
        // q.enq(neighbours...)
        // node.makeWater()
    // lvl += 1
