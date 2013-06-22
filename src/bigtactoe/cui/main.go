package main

import (
	"bigtactoe/game"
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	g := game.NewGame()
	var gameOver bool
	var winner game.Token
	for {
		clearScreen()
		gameOver, winner = g.IsGameOver()
		if gameOver {
			break
		}
		render(g)
		for {
			bigX, bigY, x, y, err := getMove(g)
			if err != nil {
				fmt.Printf("Invalid move format: %s. Please try again\n", err.Error())
				continue
			}

			fmt.Printf("{%d,%d} - {%d,%d}\n", bigX, bigY, x, y)
			err = g.PlaceToken(bigX, bigY, x, y)
			if err != nil {
				fmt.Println("Invalid move: %s. Please try again.", err.Error())
				continue
			}
			break
		}
	}
	fmt.Printf("%s is the winner!!!", getTokenName(winner))
}

func getMove(g *game.Game) (int, int, int, int, error) {
	bigX, bigY, x, y := -1, -1, -1, -1
	var err error

	if g.AnyBoardAllowed() {
		fmt.Println("Enter your board: (0 based, comma delimited)")
		bigX, bigY, err = getCoordinates()
	} else {
		bigX, bigY, err = g.GetBoardRequirementCoordinates()
	}
	if err == nil && !isValidCoordinates(bigX, bigY) {
		err = errors.New("Invalid board")
	}
	if err != nil {
		return bigX, bigY, x, y, err
	}

	fmt.Printf("Enter your move on board {%d,%d}: (0 based, comma delimited)\n", bigX, bigY)

	x, y, err = getCoordinates()
	if err == nil && !isValidCoordinates(x, y) {
		err = errors.New("Invalid coordinates")
	}
	if err != nil {
		return -1, -1, -1, -1, err
	}

	return bigX, bigY, x, y, err
}

func isValidCoordinates(x, y int) bool {
	return x >= 0 && x < 3 && y >= 0 && y < 3 // TODO Constants!
}

func getCoordinates() (int, int, error) {
	var x, y int
	_, err := fmt.Scanf("%d,%d\n", &x, &y)
	if err != nil {
		return -1, -1, err
	}
	return x, y, nil
}

func render(g *game.Game) {
	fmt.Printf("Player %s's turn\n", getTokenName(g.CurrentPlayer))
	if g.AnyBoardAllowed() {
		fmt.Println("You can move anywhere!")
	} else {
		x, y, err := g.GetBoardRequirementCoordinates()
		if err != nil {
			panic(err)
		}
		fmt.Printf("You must make your move in {%d,%d}\n", x, y)
	}
	renderGameBoard(g)
}

func renderGameBoard(g *game.Game) {
	for i := 0; i < 3; i += 1 {
		renderBoardRow(g, i)
	}
}

func renderBoardRow(g *game.Game, boardRow int) {
	startIndex, stopIndex := boardRow*3, boardRow*3+3
	boards := g.InternalBoards[startIndex:stopIndex]
	for row := 0; row < 3; row += 1 {
		for k, v := range boards {
			for col := 0; col < 3; col += 1 {
				fmt.Print(" " + getTokenName(v.Get(col, row)) + " ")
				if col == 2 && k != 2 {
					fmt.Print("|")
				}
			}
		}
		fmt.Println()
	}
	if boardRow != 2 {
		fmt.Printf("-----------------------------\n")
	}
}

func getTokenName(t game.Token) string {
	switch {
	case t == game.Blank:
		return "-"
	case t == game.X:
		return "X"
	}
	return "O"
}

func clearScreen() {
	c := exec.Command("cls")
	c.Stdout = os.Stdout
	c.Run()
}
