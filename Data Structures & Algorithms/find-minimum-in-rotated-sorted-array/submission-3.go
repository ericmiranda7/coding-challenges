func findMin(nums []int) int {
    s := 0
    e := len(nums)-1

    for s < e {
        m := ((e-s)/2)+s

        if nums[m] < nums[0] || nums[m] < nums[len(nums)-1] {
            // in left half
            e = m
        } else {
            s = m+1
        }
    }

    return nums[s]
}