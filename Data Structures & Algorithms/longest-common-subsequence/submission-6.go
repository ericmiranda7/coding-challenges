func longestCommonSubsequence(text1 string, text2 string) int {
    return solve(text1, text2, 0, 0, map[[2]int]int{})
}

func solve(text1, text2 string, p1, p2 int, cache map[[2]int]int) int {
	key := [2]int{p1, p2}
	if v, ok := cache[key]; ok {
		return v
	}

	if p1 == len(text1) || p2 == len(text2) {
		return 0
	}

	maxLcs, lcs1, lcs2 := 0, 0, 0 
	if text1[p1] == text2[p2] {
		lcs1 = 1 + solve(text1, text2, p1+1, p2+1, cache)
	} else {
		lcs1 = solve(text1, text2, p1+1, p2, cache)
		lcs2 = solve(text1, text2, p1, p2+1, cache)
	}
	maxLcs = max(lcs1, lcs2)

	cache[key] = maxLcs

	return maxLcs
}