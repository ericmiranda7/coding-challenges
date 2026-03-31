// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func mergeSort(pairs []Pair) []Pair {
    return mergeHelper(pairs, 0, len(pairs) - 1)
}

func mergeHelper(pairs []Pair, start, end int) []Pair {
    if end <= start {
        return pairs
    }

    mid := ((end - start) / 2) + start
    mergeHelper(pairs, start, mid)
    mergeHelper(pairs, mid+1, end)

    merge(pairs, start, mid, end)
    return pairs
}

func merge(pairs []Pair, start, mid, end int) {
    res := []Pair{}
    l := start
    r := mid + 1

    for l <= mid && r <= end {
        if pairs[l].Key <= pairs[r].Key {
            res = append(res, pairs[l])
            l += 1
        } else {
            res = append(res, pairs[r])
            r += 1
        }
    }

    if l > mid {
        for ;r <= end; r++ {
            res = append(res, pairs[r])
        }
    } else {
        for ;l <= mid; l++ {
            res = append(res, pairs[l])
        }
    }

    for i, v := range res {
        pairs[start+i] = v
        fmt.Println(pairs)
    }
}
