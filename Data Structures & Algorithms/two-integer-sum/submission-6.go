func twoSum(nums []int, target int) []int {
    m := map[int]int{} // val -> indx

    for j, num := range nums {
        comp := target - num
        if i, ok := m[comp]; ok {
            if j < i {
                return []int{j, i}
            } else {
                return []int{i, j}
            }
        } else {
            m[num] = j
        }
    }

    return []int{}
}
