func maxSubArray(nums []int) int {
	max := math.MinInt
    curr := 0

	sum := 0
	for curr < len(nums) {
		sum += nums[curr]
		if sum > max {
			max = sum
		}

		curr += 1
		if sum < 0 {
			sum = 0
		}
	}

	return max
}

// 2, -3, 5