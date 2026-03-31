// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
    // gotta be stable
    if len(pairs) == 0 {
        return pairs
    }
    
    return mergeSorting(pairs, 0, len(pairs) - 1 )
}

func mergeSorting(pairs []Pair, s int, e int) []Pair {
    if (e - s == 0) {
        return pairs[s:e+1]
    }

    m := ((e - s) / 2) + s
    lsorted := mergeSorting(pairs, s, m)
    rsorted := mergeSorting(pairs, m + 1, e)
    fullSorted := mergeEm(lsorted, rsorted)
    return fullSorted
}

func mergeEm(lpair []Pair, rpair []Pair) []Pair {
    res := []Pair{}

    lptr := 0
    rptr := 0

    for lptr < len(lpair) && rptr < len(rpair) {
        if lpair[lptr].Key <= rpair[rptr].Key {
            res = append(res, lpair[lptr])
            lptr += 1
        } else {
            res = append(res, rpair[rptr])
            rptr += 1
        }
    }
    if lptr >= len(lpair) {
        res = append(res, rpair[rptr:]...)
    } else {
        res = append(res, lpair[lptr:]...)
    }

    return res
}