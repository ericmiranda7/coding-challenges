func subsets(nums []int) [][]int {
    res := [][]int{}
    subset := []int{}
    helper(nums, subset, &res)
    return res
}

func helper(nums []int, subset []int, res *[][]int) {
    if len(nums) == 0 {
        cpy := make([]int, len(subset))
        copy(cpy, subset)
        *res = append(*res, cpy)
        return
    }

    pick := nums[0]
    // dont take pick
    helper(nums[1:], subset, res)
    // take pick
    subset = append(subset, pick)
    helper(nums[1:], subset, res)
}