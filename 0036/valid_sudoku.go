func isValidSudoku(board [][]byte) bool {

	for i := 0; i < 9; i++ {
		rows := make(map[byte]int)
		columns := make(map[byte]int)
		cubes := make(map[byte]int)
		for j := 0; j < 9; j++ {
			_, ok := rows[board[i][j]]
			if ok && board[i][j] != '.' {
				return false
			}
			rows[board[i][j]] ++

			_, ok = columns[board[j][i]]
			if ok && board[j][i] != '.' {
				return false
			}
			columns[board[j][i]] ++

			rowIndex := 3 * (i / 3)
			colIndex := 3 * (i % 3)

			_, ok = cubes[board[rowIndex+j/3][colIndex+j%3]]
			if ok && board[rowIndex+j/3][colIndex+j%3] != '.' {
				return false
			}
            cubes[board[rowIndex+j/3][colIndex+j%3]] ++

		}

	}
	return true

}
