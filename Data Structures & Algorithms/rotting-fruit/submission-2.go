type Node [2]int

const (
    EMPTY = 0
    FRESH = 1
    ROTTEN = 2
)

func orangesRotting(grid [][]int) int {
    rottenNodes := getRotten(grid)
    q := []Node{}
    q = append(q, rottenNodes...)
    
    lvl := -1
    for len(q) > 0 {
        qlen := len(q)
        for _ = range qlen {
            n := q[0]
            q = q[1:]

            if grid[n[0]][n[1]] == EMPTY {
                continue
            }
            if lvl != -1 && grid[n[0]][n[1]] == ROTTEN {
                continue
            }
            freshNbrs := getFreshNbrs(grid, n)
            q = append(q, freshNbrs...)
            grid[n[0]][n[1]] = EMPTY
        }
        lvl += 1
    }

    if lvl < 0 {
        lvl = 0
    }

    if containsFresh(grid) {
        return -1
    } else {
        return lvl
    }
}

func getRotten(grid [][]int) []Node {
    r := []Node{}
    for row, _ := range grid {
        for col, _ := range grid[row] {
            if grid[row][col] == ROTTEN {
                r = append(r, Node{row, col})
            }
        }
    }

    return r
}

func getFreshNbrs(grid [][]int, n Node) []Node {
    p := []Node{
        Node{n[0]-1, n[1]},
        Node{n[0], n[1]-1}, Node{n[0], n[1]+1},
        Node{n[0]+1, n[1]},
    }

    r := []Node{}
    for _, c := range p {
        if c[0] < 0 || c[0] >= len(grid) {
            continue
        }
        if c[1] < 0 || c[1] >= len(grid[c[0]]) {
            continue
        }
        if grid[c[0]][c[1]] == ROTTEN || grid[c[0]][c[1]] == EMPTY {
            continue
        }

        r = append(r, c)
    }

    return r
}

func containsFresh(grid [][]int) bool {
    for row, _ := range grid {
        for col, _ := range grid[row] {
            if grid[row][col] == FRESH {
                return true
            }
        }
    }
    return false
}
