// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
    res := [][]Pair{}
    for i, _ := range pairs {
        for j := i; j > 0 && pairs[j].Key < pairs[j-1].Key; j-- {
            // swap
            tmp := pairs[j]
            pairs[j] = pairs[j-1]
            pairs[j-1] = tmp
        }
        pcpy := make([]Pair, len(pairs))
        copy(pcpy, pairs)
        res = append(res, pcpy)
    }

    return res
}
