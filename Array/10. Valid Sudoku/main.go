package main

func isValidSudoku(board [][]byte) bool {
    var rows [9][9]bool
    var cols [9][9]bool
    var childArrs [9][9]bool
    for i := 0; i < 9; i++ {
        for j := 0; j < 9; j++ {
            item := board[i][j]
            if item == '.' {
                continue
            }
            item = item - '1'
            index := i / 3 * 3 + j / 3
            if rows[i][item] {
                return false
            }
            if cols[j][item] {
                return false
            }
            if childArrs[index][item] {
                return false
            }
            rows[i][item] = true
            cols[j][item] = true
            childArrs[index][item] = true
        }
    }
    return true
}
