package dsa

func isValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]bool

	for r := 0; r < 9; r++ {
		for c := 0; c < 9; c++ {
			if board[r][c] == '.' {
				continue
			}
			d := board[r][c] - '1'
			b := (r/3)*3 + c/3

			if row[r][d] || col[c][d] || box[b][d] {
				return false
			}
			row[r][d] = true
			col[c][d] = true
			box[b][d] = true
		}
	}
	return true
}
