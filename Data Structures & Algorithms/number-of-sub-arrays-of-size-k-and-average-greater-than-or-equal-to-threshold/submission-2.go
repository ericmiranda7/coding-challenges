func numOfSubarrays(arr []int, k int, threshold int) int {
	l := 0
	sum := 0

	r := 0
	count := 0
	for r < len(arr) {
		sum += arr[r]
		if r - l >= k {
			sum -= arr[l]
			l += 1
		}
		if r - l == k - 1 && sum / k >= threshold {
			count += 1
		}
		r += 1
	}

	return count
}
