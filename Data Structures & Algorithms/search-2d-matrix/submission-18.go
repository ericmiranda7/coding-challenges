func searchMatrix(matrix [][]int, target int) bool {
    // binsearch on 2d arr (target in arr[0]..arr[len(arr)])
    // binsearch on arr (regular)

    m := len(matrix)
    s := 0
    e := m

    for s < e {
        mid := ((e-s)/2)+s

        if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
            // binsearch arr
            arr := matrix[mid]
            start := 0
            end := len(arr)

            for start < end {
                mid := ((end - start)/2)+start
                
                if arr[mid] == target {
                    return true
                } else if target > arr[mid] {
                    start = mid+1
                } else {
                    end = mid
                }
            }
            return false
        } else if target > matrix[mid][len(matrix[mid])-1] {
            s = mid+1
        } else {
            e = mid
        }
    }

    return false
}
