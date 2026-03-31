// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
    mergeHelper(pairs, 0, len(pairs)-1)
    return pairs
}

func mergeHelper(pairs []Pair, s, e int) {
    if s >= e {
        return
    }

    m := ((e - s) / 2) + s
    mergeHelper(pairs, s, m)
    mergeHelper(pairs, m+1, e)
    merge(pairs, s, m, e)
}

func merge(pairs []Pair, s, m, e int) {
    tmp := []Pair{}
    l := s
    r := m+1

    for l <= m && r <= e {
        if pairs[l].Key <= pairs[r].Key {
            tmp = append(tmp, pairs[l])
            l+=1
        } else {
            tmp = append(tmp, pairs[r])
            r+=1
        }
    }

        tmp = append(tmp, pairs[l:m+1]...)
        tmp = append(tmp, pairs[r:e+1]...)

    for i, p := range tmp {
        pairs[i+s] = p
    }
}
