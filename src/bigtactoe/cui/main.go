package main

import (
	"bigtactoe/game"
	"fmt"
)

func main() {
	g := game.NewGame()
	var gameOver bool
	var winner game.Token
	for {
		gameOver, winner = g.IsGameOver()
		if gameOver {
			break
		}
		render(g)
		bigX, bigY, x, y, err := getMove(g)
		if err != nil {
			panic(err)
		}
		g.PlaceToken(bigX, bigY, x, y)
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
	if err != nil {
		return bigX, bigY, x, y, err
	}

	fmt.Printf("Enter your move on board {%d,%d}: (0 based, comma delimited)\n", bigX, bigY)

	x, y, err = getCoordinates()
	if err != nil {
		return -1, -1, -1, -1, err
	}
	return bigX, bigY, x, y, err
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
	for k, v := range g.InternalBoards {
		for col := 0; col < 3; col += 1 {
			for row := 0; row < 3; row += 1 {
				fmt.Print(" " + getTokenName(v.Get(col, row)) + " ")
			}
			if col != 2 {
				fmt.Print("|")
			}
		}
		fmt.Println()
		if k != 8 && (k+1)%3 == 0 {
			fmt.Printf("-----------------------------\n")
		}
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
