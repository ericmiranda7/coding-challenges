func maxSubarraySumCircular(nums []int) int {
	maxSum := nums[0]
	minSum := nums[0]

	currMax := 0
	currMin := 0

	r := 0
	totalSum := 0
	for r < len(nums) {
		totalSum += nums[r]

		currMax += nums[r]
		maxSum = max(currMax, maxSum)
		currMin += nums[r]
		minSum = min(currMin, minSum)
		r += 1

		if currMax < 0 {
			currMax = 0
		}
		if currMin > 0 {
			currMin = 0
		}
	}

	if maxSum < 0 {
		return maxSum
	}

	opt1 := maxSum
	opt2 := totalSum - minSum
	
	return max(opt1, opt2)
}
