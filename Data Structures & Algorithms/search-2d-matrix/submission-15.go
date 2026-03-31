func searchMatrix(matrix [][]int, target int) bool {
    row := 0
    col := len(matrix[0])-1

    for row < len(matrix) && col >= 0 {
        if matrix[row][col] == target {
            return true
        } else if target < matrix[row][col] {
            col--
        } else {
            row++
            col = len(matrix[0]) - 1
        }
    }
    return false
}
