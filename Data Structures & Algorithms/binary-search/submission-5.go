func search(nums []int, target int) int {
    return bsearch(nums, target, 0, len(nums) - 1)
}

func bsearch(nums []int, target, start, end int) int {
    if start > end {
        return -1
    }

    m := ((end - start) / 2) + start
    if target == nums[m] {
        return m
    } else if target > nums[m] {
        return bsearch(nums, target, m+1, end)
    } else {
        return bsearch(nums, target, start, m-1)
    }
}
