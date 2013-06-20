package game

import (
	"errors"
	"fmt"
)

type Game struct {
	CurrentPlayer, NextPlayer Token
	RequiredBoard             int
	SmallBoards               [boardSize * boardSize]Board
	BigBoard                  Board
}

func NewGame() *Game {
	g := new(Game)
	g.CurrentPlayer, g.NextPlayer = X, O
	g.BigBoard = *new(Board)
	for k, _ := range g.SmallBoards {
		g.SmallBoards[k] = *new(Board)
	}
	return g
}

// Returns Blank if the game is still running.
// Returns X or O if the games been won.
// Returns Blank if it's tied.
func (g *Game) PlaceToken(b Board, x, y int) (Token, error) {

	// Make sure it's on the right board
	bigX, bigY := g.getBoardLocation(b)
	bigIndex := coordinatesToIndex(bigX, bigY)
	if !g.AnyBoardAllowed() && g.RequiredBoard != bigIndex {
		message := fmt.Sprintf("Cannot place token {%d} board {%d,%d}", g.CurrentPlayer, x, y)
		return Blank, errors.New(message)
	}

	// Set the token
	err := b.Set(g.CurrentPlayer, x, x)
	if err != nil {
		return Blank, err
	}

	// Return the token if the game's been won
	if g.BigBoard.Get(bigX, bigY) == Blank && IsWinner(g.CurrentPlayer, g.BigBoard) {
		g.BigBoard.Set(g.CurrentPlayer, bigX, bigY)
	}
	if IsWinner(g.CurrentPlayer, g.BigBoard) {
		return g.CurrentPlayer, nil
	}

	// Return blank if the games over
	if !g.BigBoard.BlankSpacesExist() {
		return Blank, nil
	}

	// setup next turn
	g.setNextBoard(x, y)
	g.CurrentPlayer, g.NextPlayer = g.NextPlayer, g.CurrentPlayer

	return Blank, nil
}

const NoBoardRequirment int = -1

func (g *Game) AnyBoardAllowed() bool {
	return g.RequiredBoard == NoBoardRequirment
}

// Set the next required board (if it's not full)
func (g *Game) setNextBoard(x, y int) {
	b := g.getBoard(x, y)
	if b.BlankSpacesExist() {
		g.RequiredBoard = coordinatesToIndex(x, y)
	} else {
		g.RequiredBoard = NoBoardRequirment // TODO: Make constant!
	}
}

func (g *Game) getBoard(x, y int) Board {
	return g.SmallBoards[coordinatesToIndex(x, y)]
}

func (g *Game) getBoardLocation(b Board) (int, int) {
	for k, v := range g.SmallBoards {
		if v == b {
			return indexToCoordinates(k)
		}
	}
	panic("Couldn't find board!")
}
