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
    return sortHelper(pairs, 0, len(pairs) - 1)
}

func sortHelper(pairs []Pair, start, end int) []Pair {
    if end <= start {
        return pairs
    }

    pvtIndx := end
    left := start
    for i:= start ;i < pvtIndx; i++ {
        fmt.Println("ireach")
        if pairs[i].Key < pairs[pvtIndx].Key {
            pairs[left], pairs[i] = pairs[i], pairs[left]
            left += 1
        }
    }
    pairs[left], pairs[pvtIndx] = pairs[pvtIndx], pairs[left]

    sortHelper(pairs, start, left-1)
    sortHelper(pairs, left + 1, end)
    return pairs
}
