func twoSum(nums []int, target int) []int {
    hmap := map[int]int{}

    for i, v := range nums {
        hmap[v] = i
    }

    for i, v := range nums {
        if j, ok := hmap[target - v]; ok && i != j {
            return []int{i, j}
        }
    }

    return []int{}
}
