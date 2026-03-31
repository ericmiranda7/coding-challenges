func sortColors(nums []int) {
    // build a freq map color -> freq
    // itr on the "map"
        // mutate original arr[itr_indx]

    amap := [3]int{}
    for _, v := range nums {
        amap[v] += 1
    }

    ogi := 0
    for i, v := range amap {
        for _ = range v {
            nums[ogi] = i
            ogi += 1
        }
    }
}
