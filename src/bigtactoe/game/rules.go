package game

func IsWinner(t Token, b Board) bool {
	return isHorizontalWin(t, b) || isVerticalWin(t, b) || isDiagonalWin(t, b)
}

func isDiagonalWin(t Token, b Board) bool {
	return match(t, b.Get(0, 0), b.Get(1, 1), b.Get(2, 2)) ||
		match(t, b.Get(2, 0), b.Get(1, 1), b.Get(0, 2))
}

func isHorizontalWin(t Token, b Board) bool {
	for i := 0; i < boardSize; i++ {
		if match(t, b.Get(0, i), b.Get(1, i), b.Get(2, i)) {
			return true
		}
	}
	return false
}

func isVerticalWin(t Token, b Board) bool {
	for i := 0; i < boardSize; i++ {
		if match(t, b.Get(i, 0), b.Get(i, 1), b.Get(i, 2)) {
			return true
		}
	}
	return false
}

func match(tokens ...Token) bool {
	first := tokens[0]
	for _, v := range tokens[1:] {
		if first != v {
			return false
		}
	}
	return true
}
