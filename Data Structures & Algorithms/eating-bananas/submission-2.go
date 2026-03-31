func minEatingSpeed(piles []int, h int) int {
    // find min
    // find max
    // for k in min..max
        // if canEatIn(piles, k) <= h
            // return k
    // return -1
    max := piles[0]
    for _, v := range piles {
        if v > max {
            max = v
        }
    }

    start := 1
    end := max
    smallestK := 1

    for start <= end {
        m := ((end - start) / 2) + start

        hrsTaken := canEatIn(piles, m)
        if hrsTaken <= h {
            smallestK = m
            end = m - 1
        } else if hrsTaken > h {
            start = m + 1
        } else {
            end = m - 1
        }
    }

    return smallestK

}

func canEatIn(piles []int, k int) int {
    hrsTaken := 0
    for _, bananas := range piles {
        if bananas % k == 0 {
            hrsTaken += bananas / k
        } else {
            hrsTaken += (bananas / k) + 1
        }
    }
    return hrsTaken
}

// 