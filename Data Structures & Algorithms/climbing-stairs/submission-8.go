func climbStairs(n int) int {
	return climb(n, map[int]int{})
}

func climb(n int, memo map[int]int) int {
	if n <= 1 {
		return 1
	}
	if v, ok := memo[n]; ok {
		return v
	}

	twoStairs := climb(n-2, memo)
	oneStair := climb(n-1, memo)
	memo[n] = twoStairs + oneStair

	return twoStairs + oneStair
}
