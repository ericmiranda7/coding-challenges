func searchMatrix(matrix [][]int, target int) bool {
    rows := len(matrix)
    cols := len(matrix[0])

    s := 0
    e := rows*cols-1

    for s <= e {
        m := ((e-s)/2)+s

        rowIndx := (m)/cols
        colIndx := (m)%cols

        if matrix[rowIndx][colIndx] == target {
            return true
        } else if target < matrix[rowIndx][colIndx] {
            e = m-1
        } else {
            s = m+1
        }
    }
    return false
}
