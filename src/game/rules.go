package game

func getWinner(b Board) Token {
	return Blank
}

func isWinner(b Board, t Token) bool {
	return isHorizontalWin(b, t) || isVerticalWin(b, t) || isDiagonalWin(b, t)
}

func isHorizontalWin(b Board, t Token) bool {
	for i := 0; i < boardSize; i++ {
		if allMatch(t, b.Get(0, i), b.Get(1, i), b.Get(2, i)) {
			return true
		}
	}
	return false
}

func isVerticalWin(b Board, t Token) bool {
	for i := 0; i < boardSize; i++ {
		if allMatch(t, b.Get(i, 0), b.Get(i, 1), b.Get(i, 2)) {
			return true
		}
	}
	return false
}

func isDiagonalWin(b Board, t Token) bool {
	return allMatch(t, b.Get(0, 0), b.Get(1, 1), b.Get(2, 2)) ||
		allMatch(t, b.Get(2, 0), b.Get(1, 1), b.Get(0, 2))
}

func allMatch(tokens ...Token) bool {
	first := tokens[0]
	for _, v := range tokens[1:] {
		if first != v {
			return false
		}
	}
	return true
}
