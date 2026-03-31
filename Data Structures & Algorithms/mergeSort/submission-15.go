// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
    return merge(pairs, 0, len(pairs))
}

func merge(pairs []Pair, start, end int) []Pair {
    if start >= end {
        return []Pair{}
    }
    if end - start == 1 {
        return []Pair{pairs[start]}
    }

    mid := ((end - start) / 2) + start
    lhalf := merge(pairs, start, mid)
    rhalf := merge(pairs, mid, end)

    return mergeLists(lhalf, rhalf)
}

func mergeLists(lhalf, rhalf []Pair) []Pair {
    lstart := 0
    lend := len(lhalf)
    rstart := 0
    rend := len(rhalf)

    res := []Pair{}
    for lstart < lend && rstart < rend {
        if lhalf[lstart].Key <= rhalf[rstart].Key {
            res = append(res, lhalf[lstart])
            lstart++
        } else {
            res = append(res, rhalf[rstart])
            rstart++
        }
    }

    if lstart < lend {
        res = append(res, lhalf[lstart:]...)
    } else {
        res = append(res, rhalf[rstart:]...)
    }
    return res
}
