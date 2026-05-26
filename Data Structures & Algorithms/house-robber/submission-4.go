func rob(nums []int) int {
    return solve(nums, 0, map[int]int{})
}

func solve(nums []int, indx int, cache map[int]int) int {
	if v, ok := cache[indx]; ok {
		return v
	}

	if indx == len(nums) {
		return 0
	} else if indx == len(nums) - 1 {
		return nums[len(nums)-1]
	}

	// pick
	// 5 2
	pick := nums[indx] + solve(nums, indx+2, cache)

	// skip
	skip := solve(nums, indx+1, cache)

	if pick > skip {
		cache[indx] = pick
		return pick
	} else {
		cache[indx] = skip
		return skip
	}
}
