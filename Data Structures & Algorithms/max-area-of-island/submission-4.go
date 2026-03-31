type Node [2]int
const LAND = 1
const WATER = 0

func maxAreaOfIsland(grid [][]int) int {
    maxArea := 0
    for row, _ := range grid {
        for col, _ := range grid[row] {
            node := Node{row, col}
            if grid[row][col] == LAND {
                area := dfs(grid, node, 1)
                maxArea = int(math.Max(float64(maxArea), float64(area)))
            }
        }
    }

    return maxArea
}

func dfs(grid [][]int, node Node, area int) int {
    if grid[node[0]][node[1]] == WATER {
        return area - 1
    }
    grid[node[0]][node[1]] = WATER
    neighbours := getNeighbours(grid, node)
    for _, neighbour := range neighbours {
        area = dfs(grid, neighbour, area + 1)
    }
    return area
}

func getNeighbours(grid [][]int, node Node) []Node {
    neighbours := []Node{
        Node{node[0]-1, node[1]},
        Node{node[0], node[1]-1}, Node{node[0], node[1]+1},
        Node{node[0]+1, node[1]},
    }

    res := []Node{}
    for _, neighbour := range neighbours {
        if neighbour[0] < 0 || neighbour[0] >= len(grid) {
            continue
        }
        if neighbour[1] < 0 || neighbour[1] >= len(grid[neighbour[0]]) {
            continue
        }
        if grid[neighbour[0]][neighbour[1]] == WATER {
            continue
        }
        res = append(res, neighbour)
    }
    return res
}

// typedef Node [2]int

// maxArea := 0
// for each node in grid
    // if node == LAND
        // area := dfs(grid, node, 0)
        // maxArea = math.max(maxArea, area)
// return maxArea

// dfs(grid, node, area)
// grid[Node] = WATER
// neighbours := getNeighbours(grid, node)
// for each neighbour in neighbours
    // area += dfs(grid, neighbour, area + 1)
// return area

// getNeighbours(grid, node)
// neighbours := []Node{up, left, right, down}
// neighbours.filter(n -> n == 1 && (n[r], n[c] is valid))
