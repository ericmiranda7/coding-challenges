func topKFrequent(nums []int, k int) []int {
    freqMap := map[int]int{}
    for _, v := range nums {
        freqMap[v] += 1
    }

    freqGroup := make([][]int, len(nums))
    for k, v := range freqMap {
        freqGroup[v-1] = append(freqGroup[v-1], k)
    }

    res := []int{}
    for i := len(freqGroup) - 1; i >= 0; i-- {
        if len(freqGroup[i]) == 0 {
            continue
        }

        k -= len(freqGroup[i])
        res = append(res, freqGroup[i]...)
        if k == 0 {
            break
        }
    }

    return res
}
