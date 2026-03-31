func getConcatenation(nums []int) []int {
    ans := make([]int, len(nums) * 2)

    i := 0
    for i < len(nums) {
        ans[i] = nums[i]
        ans[i + len(nums)] = nums[i]
        i+=1
    }

    return ans
}
