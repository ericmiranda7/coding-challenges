func rob(nums []int) int {
    sofar := make([]int, len(nums))

	for i, _ := range nums {
		// pick
		potSoFar := getPot(sofar, i-2)
		pick := nums[i] + potSoFar

		// skip
		skip := getPot(sofar, i-1)

		if pick > skip {
			sofar[i] = pick
		} else {
			sofar[i] = skip
		}
	}

	return sofar[len(sofar)-1]
}

func getPot(sofar []int, i int) int {
	if i < len(sofar) && i >= 0 {
		return sofar[i]
	} else {
		return 0
	}
}
