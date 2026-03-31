func hasDuplicate(nums []int) bool {
    m := map[int]bool{}

    for _, num := range nums {
        if _, exists := m[num]; exists {
            return true
        } else {
            m[num] = true
        }
    }

    return false
}
