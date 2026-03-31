func hasDuplicate(nums []int) bool {
    m := map[int]bool{}

    for _, num := range nums {
        if m[num] {
            return true
        } else {
            m[num] = true
        }
    }

    return false
}
