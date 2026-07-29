func exist(board [][]byte, word string) bool {
	M := len(board)
	N := len(board[0])
	checked := make([][]bool, len(board))
	for i := 0; i < len(checked); i++ {
		checked[i] = make([]bool, len(board[0]))
	}
	var buf bytes.Buffer
	buf.Grow(len(word))
	for i := 0; i < M; i++ {
		for j := 0; j < N; j++ {
			if backtracking(board, checked, word, &buf, i, j, M, N) {
				return true
			}
		}
	}
	return false
}

func backtracking(board [][]byte, checked [][]bool, word string, buf *bytes.Buffer, row int, col int, M int, N int) bool {
	c := board[row][col]
	buf.WriteByte(c)
	currentIndex := buf.Len() - 1
	if c != word[currentIndex] {
		buf.Truncate(currentIndex) // Clean up buffer before returning
		return false
	}
	// If it matches and we reached the target length, we found the word!
	if buf.Len() == len(word) {
		return true
	}

	checked[row][col] = true
	options := [][]int{{1,0}, {-1,0}, {0,1}, {0,-1}}
	for _, opt := range options {
		newRow := row + opt[0]
		newCol := col + opt[1]
		if valid(newRow, newCol, word, buf, M, N) && !checked[newRow][newCol] {
			if backtracking(board, checked, word, buf, newRow, newCol, M, N) {
				return true
			}
		}
	}
	checked[row][col] = false
	buf.Truncate(buf.Len() - 1)
	return false
}

func valid(row int, col int, word string, buf *bytes.Buffer, M int, N int) bool {
	return row >= 0 && row < M && col >= 0 && col < N && !(buf.Len() + 1 > len(word))
}
