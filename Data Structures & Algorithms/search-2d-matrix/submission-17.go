func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    cols := len(matrix[0])
    s := 0
    e := rows*cols-1

    for s <= e {
        m := ((e-s)/2)+s

        row := m / cols
        col := m % cols
        if target == matrix[row][col] {
            return true
        } else if target < matrix[row][col] {
            e = m-1
        } else {
            s = m+1
        }
    }
    return false
}
