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

    for k := 1; k <= max; k++ {
        if canEatIn(piles, k) <= h {
            return k
        }
    }

    return -1
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