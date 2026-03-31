func sortColors(nums []int) {
    freq := [3]int{0, 0, 0}

    for _, color := range nums {
        freq[color] +=1
    }

    i := 0
    for color, freq := range freq {
        for _ = range freq {
            nums[i] = color
            i += 1
        }
    }
}
