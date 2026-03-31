func isInterleave(s1 string, s2 string, s3 string) bool {
	if len(s1) + len(s2) != len(s3) {
	 return false
	}

	return dfs(s1, s2, s3, 0, 0, map[[2]int]bool{})
}

func dfs(s1, s2, s3 string, i1, i2 int, memo map[[2]int]bool) bool {
	// base
	if res, ok := memo[[2]int{i1, i2}]; ok {
		return res
	}
	if i1 == len(s1) && i2 == len(s2) {
		return true
	}

	c1, c2 := false, false
	if i1 < len(s1) && s1[i1] == s3[i1+i2] {
		c1 = dfs(s1, s2, s3, i1+1, i2, memo)
	}
	if !c1 && i2 < len(s2) && s2[i2] == s3[i1+i2] {
		c2 = dfs(s1, s2, s3, i1, i2+1, memo)
	}
	
	memo[[2]int{i1, i2}] = c1 || c2

	return c1 || c2
}
