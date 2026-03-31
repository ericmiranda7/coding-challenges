func sortColors(nums []int) {
    farr := make([]int, 3)

    for _, v := range nums {
        farr[v] += 1
    }

    ogIndx := 0
    for color, freq := range farr {
        for i := 0; i < freq; i++ {
            nums[ogIndx] = color
            ogIndx += 1
        }
    }
}
