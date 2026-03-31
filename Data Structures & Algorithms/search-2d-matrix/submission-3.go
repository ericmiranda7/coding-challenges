func searchMatrix(matrix [][]int, target int) bool {
    return bin(matrix, target, 0, len(matrix) - 1)
}

func bin(matrix [][]int, target, start, end int) bool {
    if start > end {
        return false
    }

    m := ((end - start) / 2) + start

    if target >= matrix[m][0] && target <= matrix[m][len(matrix[m]) - 1] {
        // bin submat
        subArr := matrix[m]
        s := 0
        e := len(subArr) - 1

        for s <= e {
            altMid := ((e - s) / 2) + s
            
            if subArr[altMid] == target {
                return true
            } else if target < subArr[altMid] {
                e = altMid - 1
            } else {
                s = altMid + 1
            }
        }
        return false
    } else if target < matrix[m][0] {
        return bin(matrix, target, start, m - 1)
    } else {
        return bin(matrix, target, m+1, end)
    }
}
