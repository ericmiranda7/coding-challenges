// Definition for a pair.
// type Pair struct {
//     Key   int
//     Value string
// }

func insertionSort(pairs []Pair) [][]Pair {
    res := [][]Pair{}
    for i := 0; i < len(pairs); i++ {
        for j := i-1; j >= 0; j-- {
            if pairs[j+1].Key < pairs[j].Key {
                pairs[j+1], pairs[j] = pairs[j], pairs[j+1]
            } else {
                break
            }
        }
        cpy := make([]Pair, len(pairs))
        copy(cpy, pairs)
        res = append(res, cpy)
    }

    return res
}
