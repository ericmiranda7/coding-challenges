func combinationSum(nums []int, target int) [][]int {
    res := [][]int{}
    helper(nums, target, 0, []int{}, &res)
    return res
}

func helper(nums []int, target, i int, subset []int, res *[][]int) {
    if target < 0 || i >= len(nums) {
        return
    }
    if target == 0 {
        cpy := make([]int, len(subset))
        copy(cpy, subset)
        *res = append(*res, cpy)
        return
    }

    // dont pick
    decision := nums[i]
    helper(nums, target, i+1, subset, res)

    // pick
    subset = append(subset, decision)
    target -= decision
    helper(nums, target, i, subset, res)
}
