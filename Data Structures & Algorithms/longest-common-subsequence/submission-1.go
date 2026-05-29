func longestCommonSubsequence(text1 string, text2 string) int {
	cache := map[[2]int]int{}
    return solve(text1, text2, 0, 0, cache)
}

func solve(text1, text2 string, t1, t2 int, cache map[[2]int]int) int {
	// base case
	if t1 >= len(text1) {
		return 0
	}
	if t2 >= len(text2) {
		return 0
	}
	if v, ok := cache[[2]int{t1, t2}]; ok {
		return v
	}

	// recursive
	if text1[t1] == text2[t2] {
		pick := 1 + solve(text1, text2, t1+1, t2+1, cache)
		cache[[2]int{t1, t2}] = pick
		return pick
	}
	pick1 := solve(text1, text2, t1+1, t2, cache)
	pick2 := solve(text1, text2, t1, t2+1, cache)

	if pick1 > pick2 {
		cache[[2]int{t1,t2}]=pick1
		return pick1
	} else {
		cache[[2]int{t1,t2}]=pick2
		return pick2
	}
}
