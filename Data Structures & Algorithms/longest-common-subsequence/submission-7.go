func longestCommonSubsequence(text1 string, text2 string) int {
	i := len(text1)
	j := len(text2)
	grid := make([][]int, i+1)
	for r, _ := range grid {
		grid[r] = make([]int, j+1)
	}

	i -= 1
	for ; i >= 0; i-- {
		for j := len(grid[i])-2; j >= 0; j-- {
			if text1[i] == text2[j] {
				grid[i][j] = 1 + grid[i+1][j+1]
			} else {
				grid[i][j] = max(grid[i+1][j], grid[i][j+1])
			}
		}
	}

	return grid[0][0]
}
