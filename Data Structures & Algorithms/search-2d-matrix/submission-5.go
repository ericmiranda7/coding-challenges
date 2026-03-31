func searchMatrix(matrix [][]int, target int) bool {
    // binsearch rows with target >= row[0] && target <= row[end]
        // binsearch columns
    
    s := 0
    e := len(matrix) - 1

    for s <= e {
        m := ((e - s) / 2) + s

        row := matrix[m]
        if target >= row[0] && target <= row[len(row) - 1] {
            s := 0
            e := len(row) - 1

            for s <= e {
                m := ((e - s) / 2) + s

                if target == row[m] {
                    return true
                } else if target < row[m] {
                    e = m-1
                } else {
                    s = m+1
                }
            }
            return false
        } else if target > row[len(row) - 1] {
            s = m + 1
        } else {
            e = m - 1
        }
    }

    return false
}
