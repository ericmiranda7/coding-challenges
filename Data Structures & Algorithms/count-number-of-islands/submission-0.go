func numIslands(grid [][]byte) int {
    visited := map[[2]int]bool{}
    islands := 0
    for row, _ := range grid {
        for col, _ := range grid[row] {
            if grid[row][col] == '1' && !visited[[2]int{row, col}] {
                bfs(grid, row, col, visited)
                islands += 1
            }
        }
    }

    return islands
}

func bfs(grid [][]byte, initRow, initCol int, visited map[[2]int]bool) {
    q := [][2]int{[2]int{initRow, initCol}}

    for len(q) > 0 {
        rc := q[0]
        q = q[1:]
        row, col := rc[0], rc[1]

        // get neighbours
        neighbourIndexes := [][2]int{
            [2]int{row-1, col},
            [2]int{row, col-1}, [2]int{row, col+1},
            [2]int{row+1, col},
        }
        neighbours := [][2]int{}
        for _, nrc := range neighbourIndexes {
            nr, nc := nrc[0], nrc[1]
            if nr < 0 || nr >= len(grid) {
                continue
            }
            if nc < 0 || nc >= len(grid[nr]) {
                continue
            }
            if grid[nr][nc] == '0' {
                continue
            }
            if visited[nrc] {
                continue
            }

            neighbours = append(neighbours, [2]int{nr, nc})
        }

        q = append(q, neighbours...)
        visited[[2]int{row, col}] = true
    }
}

// var visitied map[int]bool
// var islands int
// for every row in grid
    // for every col in row
        // pos = (row * ncol) + col
        // if grid[pos] == 1 && !visited[pos]
            // bfs(grid, pos, visited)
            // islands += 1

// return islands

// bfs(grid, pos, visited)
// q := []int{pos}
// for q.IsNotEmpty
    // pos := q.deq()
    // neighbours = getNeighbours(pos)
    // neighbours.filter(n -> n == 1)
    // q.enq(neighbours...)
    // visited[pos] = true

