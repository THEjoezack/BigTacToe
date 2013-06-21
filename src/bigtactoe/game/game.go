package game

import (
	"errors"
	"fmt"
)

const noBoardRequirment int = -1

type Game struct {
	CurrentPlayer, NextPlayer Token
	RequiredBoardIndex        int // TODO: Shouldn't be exposed
	InternalBoards            [boardSize * boardSize]Board
	GameBoard                 Board
}

func NewGame() *Game {
	g := new(Game)
	g.CurrentPlayer, g.NextPlayer = X, O
	g.GameBoard = *new(Board)
	g.RequiredBoardIndex = noBoardRequirment
	for k, _ := range g.InternalBoards {
		g.InternalBoards[k] = *new(Board)
	}
	return g
}

func (g *Game) AnyBoardAllowed() bool {
	return g.RequiredBoardIndex == noBoardRequirment
}

// Places token and advances game
func (g *Game) PlaceToken(bigX, bigY, x, y int) error {

	// Make sure it's on the right board
	bigIndex := coordinatesToIndex(bigX, bigY)
	if !g.AnyBoardAllowed() && g.RequiredBoardIndex != bigIndex {
		message := fmt.Sprintf("Cannot place token {%d} on board {%d}. Required board is {%d}", g.CurrentPlayer, bigIndex, g.RequiredBoardIndex)
		return errors.New(message)
	}

	// Set the token
	b := g.getBoard(bigX, bigY)
	err := b.Set(g.CurrentPlayer, x, y)
	if err != nil {
		return err
	}

	// Mark the game board if the internal board's been newly won
	if g.GameBoard.Get(bigX, bigY) == Blank && IsWinner(g.CurrentPlayer, g.GameBoard) {
		g.GameBoard.Set(g.CurrentPlayer, bigX, bigY)
	}

	// Advance game
	g.RequiredBoardIndex = g.getNextBoardIndex(x, y)
	g.CurrentPlayer, g.NextPlayer = g.NextPlayer, g.CurrentPlayer

	return nil
}

func (g *Game) GetBoardRequirementCoordinates() (x, y int, err error) {
	if g.AnyBoardAllowed() {
		return -1, -1, errors.New("No board requirement")
	}
	x, y = indexToCoordinates(g.RequiredBoardIndex)
	return x, y, nil
}

// Returns (true,X|O) if game's been won
// Returns (true, Blank) if game is a tie
// Returns (false, Blank) if game is not over
func (g *Game) IsGameOver() (bool, Token) {
	for _, v := range [...]Token{X, O} {
		if IsWinner(v, g.GameBoard) {
			return true, v
		}
	}

	// Return blank if the games over
	if !g.GameBoard.BlankSpacesExist() {
		return true, Blank
	}
	return false, Blank
}

// Set the next required board (if it's not full)
func (g *Game) getNextBoardIndex(x, y int) int {
	b := g.getBoard(x, y)
	if b.BlankSpacesExist() {
		return coordinatesToIndex(x, y)
	}
	return noBoardRequirment
}

func (g *Game) getBoard(x, y int) *Board {
	return &g.InternalBoards[coordinatesToIndex(x, y)]
}
