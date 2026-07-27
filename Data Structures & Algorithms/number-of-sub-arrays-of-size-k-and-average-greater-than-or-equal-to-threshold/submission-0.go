func numOfSubarrays(arr []int, k int, threshold int) int {
	l, r := 0, 1

	count := 0
	sum := arr[l]
	for true {
		if r - l == k {
			if sum / k >= threshold {
				count += 1
			}

			sum -= arr[l]
			l += 1
		}

		if r == len(arr) {
			break
		}
		sum += arr[r]
		r += 1
	}

	return count
}
