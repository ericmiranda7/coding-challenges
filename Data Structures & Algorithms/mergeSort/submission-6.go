// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
    res := mergeHelper(pairs, 0, len(pairs) - 1)
    return res
}

func mergeHelper(pairs []Pair, s int, e int) []Pair {
    if s >= e {
        return pairs
    }

    m := ((e-s) / 2) + s
    mergeHelper(pairs, s, m)
    mergeHelper(pairs, m+1, e)
    merge(pairs, s, m, e)
    return pairs
}

func merge(pairs []Pair, s, m, e int) {
    lpair := []Pair{}
    lpair = append(lpair, pairs[s:m+1]...)
    rpair := []Pair{}
    rpair = append(rpair, pairs[m+1:e+1]...)

    lptr := 0
    rptr := 0
    i := s

    for lptr < len(lpair) && rptr < len(rpair) {
        if lpair[lptr].Key <= rpair[rptr].Key {
            pairs[i] = lpair[lptr]
            lptr += 1
        } else {
            pairs[i] = rpair[rptr]
            rptr += 1
        }
        i += 1
    }

    if lptr >= len(lpair) {
        for ; rptr < len(rpair); rptr++ {
            pairs[i] = rpair[rptr]
            i += 1
        }
    } else {
        for ; lptr < len(lpair); lptr++ {
            pairs[i] = lpair[lptr]
            i += 1
        }
    }
}
