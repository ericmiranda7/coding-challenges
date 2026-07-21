func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	currRow, prevRow := make([]int, len(obstacleGrid[0])+1), make([]int, len(obstacleGrid[0])+1)
	currRow[len(obstacleGrid[0])] = 1

	println("ok")

	for i := len(obstacleGrid)-1; i >= 0; i-- {
		println("oki")
		for j := len(obstacleGrid[0])-1; j >= 0; j-- {
			println("okj")
			if obstacleGrid[i][j] == 0 {
				println("okf")
				currRow[j] = prevRow[j] + currRow[j+1]
			}
		}
		prevRow = currRow
		currRow = make([]int, len(obstacleGrid[0])+1)
	}

	return prevRow[0]
}