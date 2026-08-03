func containsNearbyDuplicate(nums []int, k int) bool {
	hm := map[int]bool{}

	hm[nums[0]] = true
	l, r := 0, 1
	for r < len(nums) {
		if r - l > k {
			delete(hm, nums[l])
			l += 1
		}
		if hm[nums[r]] {
			return true
		}
		hm[nums[r]] = true
		r += 1
	}

	return false
}
