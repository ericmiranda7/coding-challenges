func search(nums []int, target int) int {
    return bin(nums, target, 0, len(nums) - 1)
}

func bin(nums []int, target, start, end int) int {
    if start > end {
        return -1
    }

    m := ((end - start) / 2) + start
    if nums[m] == target {
        return m
    } else if target > nums[m] {
        return bin(nums, target, m+1, end)
    } else {
        return bin(nums, target, start, m-1)
    }
}
