func maxSubArray(nums []int) int {
	max := math.MinInt
    for i, _ := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum > max {
				max = sum
			}
		}
	}

	return max
}
