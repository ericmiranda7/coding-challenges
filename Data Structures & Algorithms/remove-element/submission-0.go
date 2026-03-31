func removeElement(nums []int, val int) int {
    l := 0
    
    for _, ival := range nums {
        if ival != val {
            nums[l] = ival
            l += 1
        }
    }

    return l
}
