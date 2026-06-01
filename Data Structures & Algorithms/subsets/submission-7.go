func subsets(nums []int) [][]int {
	res := [][]int{}
	solve(nums, []int{}, &res)
	return res
}

func solve(nums, pick []int, res *[][]int) {
	if len(nums) == 0 {
		*res = append(*res, pick)
		return
	}

	// pick
	newPick := make([]int, len(pick))
	copy(newPick, pick)
	newPick = append(newPick, nums[0])

	nums = nums[1:]

	solve(nums, newPick, res)

	// don't pick
	solve(nums, pick, res)
}
