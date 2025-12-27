package searcha2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
    if len(matrix) == 0 || len(matrix[0]) == 0 {
        return false
    }

    m := len(matrix)
    n := len(matrix[0])

    left, right := 0, m*n-1

    for left <= right {
        mid := left + (right-left)/2

        r := mid / n
        c := mid % n

        val := matrix[r][c]

        if val == target {
            return true
        } else if val < target {
            left = mid + 1
        } else {
            right = mid - 1
        }
    }
    return false
}
