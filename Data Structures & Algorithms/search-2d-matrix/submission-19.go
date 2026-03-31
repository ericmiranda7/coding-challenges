func searchMatrix(matrix [][]int, target int) bool {
    m := len(matrix)
    n := len(matrix[0])
    end := m*n
    start := 0

    for start < end {
        mid := ((end-start)/2) + start
        
        r := mid / n
        c := mid % n

        if matrix[r][c] == target {
            return true
        } else if matrix[r][c] > target {
            end = mid
        } else {
            start = mid + 1
        }
    }

    return false
}
