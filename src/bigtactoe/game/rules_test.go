package game

import "testing"

func TestIsWinner(t *testing.T) {
	b := new(Board)
	if IsWinner(X, *b) {
		t.Errorf("Blank board should have no winner")
	}

	b = new(Board)
	b.Set(X, 0, 0)
	b.Set(X, 1, 1)
	b.Set(X, 2, 2)
	if !IsWinner(X, *b) {
		t.Errorf("X diagonal should win")
	}

	b = new(Board)
	b.Set(X, 0, 0)
	b.Set(X, 1, 0)
	b.Set(X, 2, 0)
	if !IsWinner(X, *b) {
		t.Errorf("Horizontal should win")
	}

	b = new(Board)
	b.Set(O, 0, 0)
	b.Set(O, 0, 1)
	b.Set(O, 0, 2)
	if !IsWinner(O, *b) {
		t.Errorf("Vertical should win")
	}
}

func TestMatch(t *testing.T) {
	if !match(X, X, X, X) {
		t.Errorf("X should all match")
	}
	if !match(O, O, O, O) {
		t.Errorf("O should all match")
	}
	if !match(Blank, Blank, Blank, Blank) {
		t.Errorf("Blank should all match")
	}
	if match(X, O, O, O) {
		t.Errorf("Mismatch shouldn't count")
	}
	if match(X, O, O, Blank) {
		t.Errorf("Mismatch w/ blank shouldn't count")
	}
}

func TestIsDiagonalWin(t *testing.T) {
	b := new(Board)
	b.Set(X, 0, 0)
	b.Set(X, 1, 1)
	b.Set(X, 2, 2)
	if !isDiagonalWin(X, *b) {
		t.Errorf("X diagonal should win")
	}

	b = new(Board)
	b.Set(O, 2, 0)
	b.Set(O, 1, 1)
	b.Set(O, 0, 2)
	if !isDiagonalWin(O, *b) {
		t.Errorf("O diagonal should win")
	}
}

func TestIsHorizontalWin(t *testing.T) {
	b := new(Board)
	b.Set(X, 0, 0)
	b.Set(X, 1, 0)
	b.Set(X, 2, 0)
	if !isHorizontalWin(X, *b) {
		t.Errorf("Horizontal should win")
	}

	b = new(Board)
	b.Set(O, 0, 0)
	b.Set(O, 0, 1)
	b.Set(O, 0, 2)
	if isHorizontalWin(O, *b) {
		t.Errorf("Vertical should not win")
	}
}

func TestIsVerticalWin(t *testing.T) {
	b := new(Board)
	b.Set(X, 0, 0)
	b.Set(X, 1, 0)
	b.Set(X, 2, 0)
	if isVerticalWin(X, *b) {
		t.Errorf("Horizontal should not win")
	}

	b = new(Board)
	b.Set(O, 0, 0)
	b.Set(O, 0, 1)
	b.Set(O, 0, 2)
	if !isVerticalWin(O, *b) {
		t.Errorf("Vertical should win")
	}
}
