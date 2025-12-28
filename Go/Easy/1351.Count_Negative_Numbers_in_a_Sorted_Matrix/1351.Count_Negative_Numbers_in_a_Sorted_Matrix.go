package countnegativenumbersinasortedmatrix

func countNegatives(grid [][]int) int {
    if len(grid) == 0 {
        return 0
    }
    c := 0
    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[i]); j++ {
            if grid[i][j] < 0 {
                c++
            }
        }
    }
    return c
}
