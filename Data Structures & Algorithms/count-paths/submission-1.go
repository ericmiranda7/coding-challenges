func uniquePaths(m int, n int) int {
	cache := map[[2]int]int{}
    return findPaths(0, 0, m, n, cache)
}

func findPaths(r, c, rows, cols int, cache map[[2]int]int) int {
	if v, ok := cache[[2]int{r,c}]; ok {
		return v
	}
	if r == rows-1 && c == cols - 1 {
		return 1
	}
	if r == rows || c == cols {
		return 0
	}

	downPaths := findPaths(r+1, c, rows, cols, cache)
	rightPaths := findPaths(r, c+1, rows, cols, cache)

	res := downPaths + rightPaths
	cache[[2]int{r,c}] = res
	return res
}
