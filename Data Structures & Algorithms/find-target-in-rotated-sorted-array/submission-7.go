func search(nums []int, target int) int {
    s := 0
    e := len(nums)-1

    for s <= e {
        m := ((e-s)/2)+s

        if target == nums[m] {
            return m
        } else if nums[m] < nums[0] || nums[m] < nums[len(nums)-1] {
            // og left half
            // graph right
            if target < nums[m] || target > nums[len(nums)-1] {
                e = m-1
            } else {
                s = m+1
            }
        } else {
            // og right half
            // graph left
            if target > nums[m] || target < nums[0] {
                s = m+1
            } else {
                e = m-1
            }
        }
    }

    return -1
}
