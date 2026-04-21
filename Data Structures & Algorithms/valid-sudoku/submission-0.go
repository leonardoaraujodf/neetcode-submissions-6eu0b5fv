func isValidSudoku(board [][]byte) bool {
	// 1. We need 9 maps to account for each row
	// 2. We need 9 maps to account for each col
	// 3. We need 9 maps to account for each sub-box
	// 4. We first need to populate these maps
	// 5. Then we iterate through each maps checking for occurrences. If there is
	// more than one occurence, sudoku is invalid. If once one occurrence for each number,
	// sudoku is valid.
	rowsMaps := make([]map[byte]byte, len(board))
	colsMaps := make([]map[byte]byte, len(board))
	boxMaps  := make([]map[byte]byte, len(board))
	rows := len(board)
	cols := len(board[0])

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
            if (board[row][col] == '.') {
                continue
            }

			if (rowsMaps[row] == nil) {
                rowsMaps[row] = make(map[byte]byte)
            }
            rowsMaps[row][board[row][col]]++
			if rowsMaps[row][board[row][col]] > 1 {
                return false
            }

            if (colsMaps[col] == nil) {
                colsMaps[col] = make(map[byte]byte)
            }
            colsMaps[col][board[row][col]]++
            if colsMaps[col][board[row][col]] > 1 {
                return false
            }
			
            if (boxMaps[(row / 3) * 3 + col / 3] == nil) {
                boxMaps[(row / 3) * 3 + col / 3] = make(map[byte]byte)
            }
            boxMaps[(row / 3) * 3 + col / 3][board[row][col]]++
			if boxMaps[(row / 3) * 3 + col / 3][board[row][col]] > 1 {
                return false
            }
            
            // fmt.Printf("boardMap %d\n", (row / 3) * 3 + col / 3)
		}
	}

	return true
}