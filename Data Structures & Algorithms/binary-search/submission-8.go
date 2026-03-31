func search(nums []int, target int) int {
    s := 0
    e := len(nums)-1

    for s <= e {
        m := ((e-s) / 2)+s

        if nums[m] == target {
            return m
        } else if target < nums[m] {
            e = m-1
        } else {
            s= m+1
        }
    }

    return -1
}
