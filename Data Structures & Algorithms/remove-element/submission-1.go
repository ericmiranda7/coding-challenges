func removeElement(nums []int, val int) int {
    // l, r = 0
    // for r < len(nums)
        // if nums[r] != val
            // nums[l] = nums[r]
            // l += 1
        // r += 1

    // return l

    l, r := 0, 0
    for r < len(nums) {
        if nums[r] != val {
            nums[l] = nums[r]
            l += 1
        }
        r += 1
    }

    return l
}
