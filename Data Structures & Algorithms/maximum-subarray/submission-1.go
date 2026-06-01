func maxSubArray(nums []int) int {
	max := nums[0]
	currSum := nums[0]
    for i := 1; i < len(nums); i++ {
		if currSum >= 0 {
			currSum += nums[i]
		} else {
			// currSum is negative, discard
			currSum = nums[i]
		}

		if currSum > max {
			max = currSum
		}
	}
	return max
}
