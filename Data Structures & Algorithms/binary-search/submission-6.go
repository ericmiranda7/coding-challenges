func search(nums []int, target int) int {
    return binsearch(nums, target, 0, len(nums) - 1)
}

func binsearch(nums []int, target, s, e int) int {
    if s > e {
        return -1
    }

    m := ((e - s) / 2)+s
    if nums[m] == target {
        return m
    } else if target < nums[m] {
        return binsearch(nums, target, s, m-1)
    } else {
        return binsearch(nums, target, m+1, e)
    }
}
