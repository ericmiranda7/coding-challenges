func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums) * 2)

    i := 0
    for i < len(nums) * 2 {
        ans[i] = nums[i % len(nums)]
        i+=1
    }

    return ans
}
