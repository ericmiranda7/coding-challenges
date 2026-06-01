func longestCommonSubsequence(text1 string, text2 string) int {
	grid := make([][]int, len(text1)+1)
	for i := range len(text1) + 1 {
		grid[i] = make([]int, len(text2)+1)
	}


	for t1 := len(text1) - 1; t1 >= 0; t1-- {
		for t2 := len(text2) - 1; t2 >= 0; t2-- {

			if text1[t1] == text2[t2] {
				grid[t1][t2] = 1 + grid[t1+1][t2+1]
			} else {
				grid[t1][t2] = int(math.Max(float64(grid[t1+1][t2]), float64(grid[t1][t2+1])))
			}
		}
	}

	return grid[0][0]
}