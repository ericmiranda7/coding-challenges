func getConcatenation(nums []int) []int {
    ans := make([]int, 2*len(nums))

    for i, _ := range ans {
        ans[i] = nums[i % len(nums)]
    }

    return ans
}
