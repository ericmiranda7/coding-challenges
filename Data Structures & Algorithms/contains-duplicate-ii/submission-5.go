func containsNearbyDuplicate(nums []int, k int) bool {
	l, r := 0, 1
	set := map[int]bool{}
	set[nums[l]] = true

	for r < len(nums) {
		if r - l > k {
			set[nums[l]] = false
			l += 1
		}
		if set[nums[r]] {
			return true
		}

		set[nums[r]] = true
		r += 1
	}

	return false
}
