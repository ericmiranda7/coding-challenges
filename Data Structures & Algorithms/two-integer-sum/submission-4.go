func twoSum(nums []int, target int) []int {
    hmap := map[int]int{}

    for i, v := range nums {
        comp := target - v
        if j, ok := hmap[comp]; ok {
            return []int{j, i}
        } else {
            hmap[v] = i
        }
    }

    return []int{}
}