func topKFrequent(nums []int, k int) []int {
    freqMap := map[int]int{}

    for _, v := range nums {
        freqMap[v] += 1
    }

    fmt.Println("hi")
    freqGroup := make([][]int, len(nums))
    fmt.Println("hi2")
    for k, v := range freqMap {
        fmt.Println(k, " ", v)
        freqGroup[v-1] = append(freqGroup[v-1], k)
        fmt.Println("um?")
    }

    fmt.Println("bye")
    fmt.Println(freqGroup)
    fmt.Println("why")

    res := []int{}
    for k > 0 {
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
    }

    return res
}
