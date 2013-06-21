package game

import "testing"

func TestNewGame(t *testing.T) {
	g := NewGame()
	if g.CurrentPlayer != X {
		t.Errorf("X should be the starting player")
	}
	if g.NextPlayer != O {
		t.Errorf("O should be the next player")
	}
	if !g.GameBoard.BlankSpacesExist() {
		t.Errorf("Game board should be initialized")
	}
	if !g.InternalBoards[0].BlankSpacesExist() {
		t.Errorf("Internal boards should be initialized")
	}
}

func TestAnyBoardAllowed(t *testing.T) {
	g := NewGame()
	if !g.AnyBoardAllowed() {
		t.Errorf("Any board should be allowed at the beginning of the game")
	}
	g.PlaceToken(1, 2, 1, 1)
	if g.AnyBoardAllowed() {
		t.Errorf("A board is required if a move is available")
	}
}

func TestIsGameOver(t *testing.T) {
	g := NewGame()
	gameOver, _ := g.IsGameOver()
	if gameOver {
		t.Errorf("Blank board should not have winner")
	}

	g = NewGame()
	g.GameBoard.Set(X, 0, 0)
	g.GameBoard.Set(X, 1, 0)
	g.GameBoard.Set(X, 2, 0)
	gameOver, winner := g.IsGameOver()
	if !gameOver || winner != X {
		t.Errorf("X should have won")
	}

	g = NewGame()
	g.GameBoard.Set(X, 0, 0)
	g.GameBoard.Set(O, 1, 0)
	g.GameBoard.Set(X, 2, 0)
	g.GameBoard.Set(X, 0, 1)
	g.GameBoard.Set(O, 1, 1)
	g.GameBoard.Set(O, 2, 1)
	g.GameBoard.Set(O, 0, 2)
	g.GameBoard.Set(X, 1, 2)
	g.GameBoard.Set(X, 2, 2)
	gameOver, winner = g.IsGameOver()
	if !gameOver || winner != Blank {
		t.Errorf("Should be tie game")
	}
}

func TestPlaceToken(t *testing.T) {
	g := NewGame()

	// X goes first, can go anywhere
	err := g.PlaceToken(0, 0, 1, 1)
	if err != nil {
		t.Error(err)
	}

	if g.CurrentPlayer != O {
		t.Errorf("Current player should now be O")
	}

	if g.RequiredBoardIndex != 4 {
		t.Errorf("Required board should have been set")
	}

	if X != g.InternalBoards[0].Get(1, 1) {
		t.Errorf("X was not retrieved: %d", g.InternalBoards[0].Get(1, 1))
	}

	// Try to make an invalid move
	err = g.PlaceToken(0, 0, 2, 2)
	if err == nil {
		t.Errorf("Placing an invalid token should have return err")
	}
	if g.RequiredBoardIndex != 4 {
		t.Errorf("Required board should still be set")
	}

	// Try a valid move
	err = g.PlaceToken(1, 1, 1, 1)
	if err != nil {
		t.Error(err)
	}
}

func TestGetNextBoardIndex(t *testing.T) {
	g := NewGame()
	if 4 != g.getNextBoardIndex(1, 1) {
		t.Errorf("{1,1} should be required")
	}
	b := g.getBoard(1, 1)
	b.Set(X, 0, 0)
	b.Set(X, 0, 1)
	b.Set(X, 0, 2)
	b.Set(X, 1, 0)
	b.Set(X, 1, 1)
	b.Set(X, 1, 2)
	b.Set(X, 2, 0)
	b.Set(X, 2, 1)
	b.Set(X, 2, 2)
	if noBoardRequirment != g.getNextBoardIndex(1, 1) {
		t.Errorf("No board should be required if the required board is totally full")
	}
}

func TestGetBoard(t *testing.T) {
	g := NewGame()
	b := g.getBoard(1, 1)
	b.Set(X, 1, 1)
	if g.InternalBoards[4].Get(1, 1) != X {
		t.Errorf("Correct board was not retrieved")
	}
}
