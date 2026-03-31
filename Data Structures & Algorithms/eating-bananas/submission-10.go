func minEatingSpeed(piles []int, h int) int {
    maxK := 0
    for _, qty := range piles {
        if qty > maxK {
            maxK = qty
        }
    }

    s := 1
    e := maxK

    ans := 0
    for s <= e {
        m := ((e-s) / 2)+s

        if canEatIn(piles, m, h) {
            ans = m
            e = m-1
        } else {
            s = m+1
        }
    }

    return ans
}

// canEatIn returns whether or not all bananas in pile can be eaten in h hours at k bananas per hour
func canEatIn(piles []int, k, h int) bool {
    hrsTaken := 0
    for _, bananas := range piles {
        hrsTaken += bananas / k
        if bananas % k != 0 {
            hrsTaken+=1
        }
    }
    return hrsTaken <= h
}
