// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

type Solution struct {
    
}

func NewSolution() *Solution {
    return &Solution{}
}

func QuickSort(pairs []Pair) []Pair {
    return quickHelper(pairs, 0, len(pairs) - 1)
}

func quickHelper(pairs []Pair, s, e int) []Pair {
    if s >= e {
        return pairs
    }

    l := s
    i := s
    for i < e {
        if pairs[i].Key < pairs[e].Key {
            tmp := pairs[l]
            pairs[l] = pairs[i]
            pairs[i] = tmp
            l+=1
        }
        i+=1
    }
    tmp := pairs[l]
    pairs[l] = pairs[e]
    pairs[e] = tmp

    quickHelper(pairs, s, l-1)
    quickHelper(pairs, l + 1, e)
    return pairs
}
