func uniquePaths(m int, n int) int {
	prevRow, currRow := make([]int, n+1), make([]int, n+1)
	currRow[n] = 1
	
	for i := 0; i < m; i++ {
		for j := n-1; j >= 0; j-- {
			currRow[j] = prevRow[j] + currRow[j+1]
		}
		prevRow = currRow
		currRow = make([]int, n+1)
	}

	return prevRow[0]
}