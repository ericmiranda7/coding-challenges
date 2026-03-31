func sortColors(nums []int) {
    // make intarr of 3 (0,1,2)
    // for color in inp arr
        // intarr[color]++
    // opptr = 0
    // for color in intarr
        // for freq in intarr[i]
            // nums[opptr] = color
            // opptr += 1

    intarr := [3]int{}
    for _, color := range nums {
        intarr[color] += 1
    }

    opptr := 0
    for color, freq := range intarr {
        for _ = range(freq) {
            nums[opptr] = color
            opptr += 1
        }
    }
}
