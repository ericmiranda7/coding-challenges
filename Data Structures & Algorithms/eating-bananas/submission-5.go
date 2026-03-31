func minEatingSpeed(piles []int, h int) int {
    maxVal := 0
    for _, v := range piles {
        if v > maxVal {
            maxVal = v
        }
    }

    l := 1
    e := maxVal

    for l < e {
        m := ((e - l) / 2)+l

        if canEatIn(piles, m) <= h {
            e = m
        } else {
            l = m+1
        }
    }

    return l

}

// returns hours in which bonanzas are eaten
func canEatIn(piles []int, eatRate int) int {
    h := 0
    for _, v := range piles {
        h += (v-1) / eatRate
        h++
    }
    return h
}
