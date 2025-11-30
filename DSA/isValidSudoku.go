func isValidSudoku(board [][]byte) bool {
	rows := make(map[string]bool, 9)
	cols := make(map[string]bool, 9)
	boxes := make(map[string]bool, 9)

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			row := fmt.Sprintf("%c-%d-%c", "R", i, board[i][j])
			col := fmt.Sprintf("%c-%d-%c", "C", j, board[i][j])

			index := (i/3)*3 + j/3
			box := fmt.Sprintf("%c-%d-%c", "B", index, board[i][j])

			if rows[row] || cols[col] || boxes[box] {
				return false
			}

			rows[row] = true
			cols[col] = true
			boxes[box] = true
		}
	}
	return true
}