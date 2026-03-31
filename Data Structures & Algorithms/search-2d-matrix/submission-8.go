func searchMatrix(matrix [][]int, target int) bool {
    // 10
    // 1, 0
    // row      col
    // itr / n, itr % n

    s := 0
    noRows := len(matrix) // 3
    noCols := len(matrix[0]) // 4
    e := (noRows * noCols) - 1

    for s <= e {
        // s=3 e=4
        m := ((e - s) / 2) + s
        // m=3

        row := m / noCols
        col := m % noCols
        // r=, c=1
        if target == matrix[row][col] {
            return true
        } else if target > matrix[row][col] {
            s = m+1
        } else {
            e = m-1
        }
    }

    return false
}
