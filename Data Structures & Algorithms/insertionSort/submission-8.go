// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
    res := [][]Pair{}
    for i := 0; i < len(pairs); i++ {
        for j := i; j > 0 && pairs[j].Key < pairs[j-1].Key; j-- {
            pairs[j], pairs[j-1] = pairs[j-1], pairs[j]
        }
        cpy := make([]Pair, len(pairs))
        copy(cpy, pairs)
        res = append(res, cpy)
    }

    return res
}
