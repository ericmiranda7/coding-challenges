func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	if len(obstacleGrid) == 1 && len(obstacleGrid[0]) == 1 && obstacleGrid[0][0] == 1{
		return 0
	}
	return solve(obstacleGrid, 0, 0, map[[2]int]int{})	
}

func solve(obstacleGrid [][]int, r, c int, cache map[[2]int]int) int {
	if v, ok := cache[[2]int{r,c}]; ok {
		return v
	}
	if r == len(obstacleGrid) - 1 && c == len(obstacleGrid[r]) - 1 {
		return 1
	}
	if r >= len(obstacleGrid) || c >= len(obstacleGrid[r]) {
		return 0
	}
	if r < len(obstacleGrid) && c < len(obstacleGrid[r]) && obstacleGrid[r][c] == 1 {
		return 0
	}


	// down
	down := solve(obstacleGrid, r+1, c, cache)

	// right
	right := solve(obstacleGrid, r, c+1, cache)

	res := down+right
	
	cache[[2]int{r,c}] = res
	return res
}
