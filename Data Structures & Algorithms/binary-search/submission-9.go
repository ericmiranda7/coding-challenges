func search(nums []int, target int) int {
    s := 0
    e := len(nums)

    for s < e {
        m := ((e - s) / 2) + s
        
        if nums[m] == target {
            return m
        } else if target > nums[m] {
            s = m+1
        } else {
            e = m
        }
    }

    return -1
}
