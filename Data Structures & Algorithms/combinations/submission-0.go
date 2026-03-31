func combine(n int, k int) [][]int {
    res := [][]int{}
    helper(n, k, 1, []int{}, &res)
    return res
}

func helper(n, k, i int, subset []int, res *[][]int) {
    if len(subset) == k {
        cpy := make([]int, len(subset))
        copy(cpy, subset)
        *res = append(*res, cpy)
        return
    }
    if i > n {
        return
    }

    for ; i <= n; i++ {
        subset = append(subset, i)
        helper(n, k, i+1, subset, res)
        subset = subset[:len(subset) - 1]
    }
}
