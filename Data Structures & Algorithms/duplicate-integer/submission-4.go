func hasDuplicate(nums []int) bool {
    set := map[int]bool{}
    for _, v := range nums {
        if set[v] {
            return true
        } else {
            set[v] = true
        }
    }

    return false
}
