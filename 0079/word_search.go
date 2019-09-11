func exist(board [][]byte, word string) bool {
	b := []byte(word)
	for x := range board {
		for y := range board[x] {
			if btracking(board, x, y, b) {
				return true
			}
		}
	}
	return false

}

func btracking(board [][]byte, x, y int, word []byte) bool {
	if x >= len(board) || x < 0 {
		return false
	}
	if y >= len(board[x]) || y < 0 {
		return false
	}
	if board[x][y] != word[0] {
		return false
	}
	if len(word) == 1 {
		return true
	}

	tmp := board[x][y]
	board[x][y] = '-'
	ret := btracking(board, x-1, y, word[1:]) ||
		btracking(board, x+1, y, word[1:]) ||
		btracking(board, x, y-1, word[1:]) ||
		btracking(board, x, y+1, word[1:])
	board[x][y] = tmp
	return ret
}
