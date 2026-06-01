func maxSubArray(nums []int) int {
	max := math.MinInt
    for i := range len(nums) {
		for j := i+1; j <= len(nums); j++ {
			sum := getSum(nums[i:j])
			if sum > max {
				max = sum
			}
		}
	}
	return max
}

func getSum(nums []int) int {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	return sum
}
