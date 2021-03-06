package game

import "testing"

// Note: Don't mind testing more than one thing per func in go
// since the test doesn't bail when there's a failure!

// Test BlankSpacesExist

func TestBlankSpacesExist(t *testing.T) {
	b := new(Board)
	if !b.BlankSpacesExist() {
		t.Errorf("New board should have blank spaces")
	}

	b.matrix[1] = X
	if !b.BlankSpacesExist() {
		t.Errorf("Board with spaces should have spaces")
	}

	for k, _ := range b.matrix {
		b.matrix[k] = O
	}

	if b.BlankSpacesExist() {
		t.Errorf("Full board should not have spaces")
	}
}

// Test Get

func TestGet(t *testing.T) {
	b := new(Board)
	checkGet(b, 2, 2, Blank, t)

	b.matrix[8] = X
	checkGet(b, 2, 2, X, t)

	b.matrix[8] = O
	checkGet(b, 2, 2, O, t)
}

func checkGet(b *Board, x, y int, expected Token, t *testing.T) {
	actual := b.Get(x, y)
	if actual != expected {
		t.Errorf("{%d,%d} should have returned %d: %d", x, y, expected, actual)
	}
}

// Test Set

func TestSet(t *testing.T) {
	b := new(Board)
	checkSet(b, X, 0, 0, t)
	checkSet(b, X, 1, 1, t)
	checkSet(b, O, 2, 2, t)
}

func TestInvalidSet(t *testing.T) {
	b := new(Board)

	err := b.Set(Blank, 0, 0)
	if err == nil {
		t.Errorf("Expected error setting blank set")
	}

	b.Set(X, 0, 0)
	err = b.Set(X, 0, 0)
	if err == nil {
		t.Errorf("Expected error setting token in used space")
	}
}

func checkSet(b *Board, expected Token, x, y int, t *testing.T) {
	err := b.Set(expected, x, y)
	if err != nil {
		t.Error(err)
		return
	}
	actual := b.matrix[coordinatesToIndex(x, y)] // Cheating!
	if actual != expected {
		t.Errorf("{%d,%d} should have returned %d: %d", x, y, expected, actual)
	}
}

//Test coordinatesToIndex

func TestCoordinatesToIndex(t *testing.T) {
	checkCoordinatesToIndex(0, 0, 0, t)
	checkCoordinatesToIndex(3, 0, 1, t)
	checkCoordinatesToIndex(5, 2, 1, t)
	checkCoordinatesToIndex(4, 1, 1, t)
	checkCoordinatesToIndex(8, 2, 2, t)
}

func checkCoordinatesToIndex(expected, x, y int, t *testing.T) {
	actual := coordinatesToIndex(x, y)
	if actual != expected {
		t.Errorf("{%d,%d} should evaluate to {%d}: {%d}", x, y, expected, actual)
	}
}

//Test indexToCoordinates

func TestIndexToCoordinates(t *testing.T) {
	checkIndexToCoordinates(0, 0, 0, t)
	checkIndexToCoordinates(1, 0, 1, t)
	checkIndexToCoordinates(2, 1, 5, t)
	checkIndexToCoordinates(1, 1, 4, t)
	checkIndexToCoordinates(2, 2, 8, t)
}

func checkIndexToCoordinates(expectedX, expectedY, input int, t *testing.T) {
	actualX, actualY := indexToCoordinates(input)
	if actualX != expectedX || actualY != expectedY {
		t.Errorf("{%d} should evaluate to {%d,%d} : {%d,%d}", input, expectedX, expectedY, actualX, actualY)
	}
}
