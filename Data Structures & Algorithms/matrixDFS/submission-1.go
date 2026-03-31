func countPaths(grid [][]int) int {
    if grid[0][0] == 1 {
        return 0
    }
    return dfs(grid, 0, map[int]bool{})
}


func dfs(grid [][]int, currentNode int, visited map[int]bool) int {
    // consts
    LAST_NODE := (len(grid) * len(grid[0])) - 1
    visited[currentNode] = true

    if currentNode == LAST_NODE {
        delete(visited, currentNode)
        return 1
    }
    if currentNode > LAST_NODE {
        delete(visited, currentNode)
        return 0
    }

    // get options
    r := currentNode / len(grid[0])
    c := currentNode % len(grid[0])
    possibilities := [][]int{
        []int{r-1,c},
        []int{r, c-1}, []int{r, c+1},
        []int{r+1, c},
    }
    realities := []int{}
    for _, poss := range possibilities {
        r := poss[0]
        c := poss[1]
        if r < 0 || r >= len(grid) {
            continue
        }
        if c < 0 || c >= len(grid[0]) {
            continue
        }

        nn := poss[0] * len(grid[0]) + poss[1]
        if grid[r][c] == 0 && !visited[nn] {
            realities = append(realities, nn)
        }
    }

    // if options == empty
        // return 0
    if len(realities) == 0 {
        delete(visited, currentNode)
        return 0
    }

    // for each option
        // s += dfs()
    uniquePaths := 0
    for _, option := range realities {
        uniquePaths += dfs(grid, option, visited)
    }

    delete(visited, currentNode)
    return uniquePaths
}


// uniquely identify r,c by single number
    // (row * ncols) + col
// reverse it
    // col : pos % ncol
    // row : pos / ncol