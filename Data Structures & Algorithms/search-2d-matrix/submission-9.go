func searchMatrix(matrix [][]int, target int) bool {
    s := 0
    e := len(matrix) - 1

    for s <= e {
        m := ((e-s) / 2)+s

        row := matrix[m]
        if target >= row[0] && target <= row[len(row)-1] {
            s := 0
            e := len(row)-1

            for s <= e {
                m := ((e-s)/2)+s

                if row[m] == target {
                    return true
                } else if target < row[m] {
                    e = m-1
                } else {
                    s = m+1
                }
            }
            return false
        } else if target < row[0] {
            e = m-1
        } else {
            s = m+1
        }
    }
    return false
}
