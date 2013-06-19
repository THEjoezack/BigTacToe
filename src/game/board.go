package game

import (
	"errors"
	"fmt"
)

type Token int

const boardSize int = 3
const (
	Blank Token = iota
	X
	O
)

type Board struct {
	matrix [boardSize * boardSize]Token
}

func (b *Board) Get(x, y int) (Token, error) {
	i := coordinatesToIndex(x, y)
	if i < 0 || i >= len(b.matrix) {
		message := fmt.Sprintf("Unable to set value, invalid index requested: %d,%d", x, y)
		return Blank, errors.New(message)
	}
	return b.matrix[i], nil
}

func (b *Board) Set(t Token, x, y int) error {
	if t == Blank {
		message := fmt.Sprintf("Unable to set blank token at {%d,%d}", x, y)
		return errors.New(message)
	}

	current, err := b.Get(x, y)
	if err != nil {
		return err
	}

	if current != Blank {
		message := fmt.Sprintf("Unable to set location {%d,%d}, space is occupied by %d", x, y, current)
		return errors.New(message)
	}
	i := coordinatesToIndex(x, y)
	b.matrix[i] = t
	return nil
}

func coordinatesToIndex(x, y int) int {
	return x + boardSize*y
}
