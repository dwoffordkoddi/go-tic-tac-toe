package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Board [][]string

type Player struct {
	number int
	character string
}

var players = []Player{{number: 1, character: "X"}, {number: 2, character: "O"}}
var currentPlayer Player
var board =	Board{
	{"-", "-", "-"},
	{"-", "-", "-"},
	{"-", "-", "-"},
}

func startGame() {
	currentPlayer = players[0]
	printBoard()
	getPlayerSlot(currentPlayer)
}

func getPlayerSlot(currentPlayer Player) {
	var allowedValues = []string{"A1", "A2", "A3", "B1", "B2", "B3", "C1", "C2", "C3"}

	var slot string
	fmt.Print(fmt.Sprintf("Player %d please select a slot: ", currentPlayer.number))
	_, err := fmt.Scanln(&slot)

	if err != nil {
		handleError(err)
	}

	if len(slot) > 2 {
		for len(slot) > 2 {
			fmt.Print("You've entered more characters than are allowed, please try again: ")
			_, err := fmt.Scanln(&slot)

			if err != nil {
				handleError(err)
			}
		}
	}

	if !contains(allowedValues, slot) {
		for !contains(allowedValues, slot) {
			fmt.Println("You've selected a slot that does not exist, please try again: ")
			_, err := fmt.Scanln(&slot)

			if err != nil {
				handleError(err)
			}
		}

	}

	var column = getColumn(slot)
	var row = getRow(slot)

	if board[column][row] == "-" {
		board[column][row] = currentPlayer.character
	} else {
		alreadySelected := true

		for alreadySelected {
			fmt.Print("You've selected a slot that has already been taken, please try again: ")
			_, err := fmt.Scanln(&slot)

			if err != nil {
				handleError(err)
			}

			column = getColumn(slot)
			row = getRow(slot)

			if board[column][row] == "-" {
				board[column][row] = currentPlayer.character
				alreadySelected = false
			} else {
				alreadySelected = true
			}
		}
	}


	printBoard()
	checkWin(currentPlayer)

	if currentPlayer.number == 1 {
		currentPlayer = players[1]
	} else {
		currentPlayer = players[0]
	}

	getPlayerSlot(currentPlayer)
}

func printBoard() {
	var slot string
	var boardLayout = `
     A   B   C
  1  A0| B0| C0
    ---|---|---
  2  A1| B1| C1
    ---|---|---
  3  A2| B2| C2 
`
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			switch i {
				case 0:
					slot = "A"
				case 1:
					slot = "B"
				case 2:
					slot = "C"
			}
			boardLayout = strings.Replace(boardLayout, slot + strconv.Itoa(j), board[i][j] + " ", 1)
		}
	}
	fmt.Print(boardLayout)
}

func contains(slice []string, string string) bool {
	for _, a := range slice {
		if a == string {
			return true
		}
	}
	return false
}

func handleError(err error) {
	fmt.Println("An error has occurred: " + err.Error())
	os.Exit(1)
}

func getColumn(slot string) int {
	var inputColumn = slot[:1]
	switch inputColumn {
		case "A":
			return 0
		case "B":
			return 1
		case "C":
			return 2
	}

	panic("Not sure how, but you've managed to get this far with a bad column selection")
}

func getRow(slot string) int {
	var inputRow = slot[1:]
	row, err := strconv.Atoi(inputRow)

	if err != nil {
		handleError(err)
	}

	return row - 1
}

func checkWin(currentPlayer Player) {
	win := false
	character := currentPlayer.character

	// Check the columns
	if board[0][0] == character && board[0][1] == character && board[0][2] == character {
		win = true
	}

	if board[1][0] == character && board[1][1] == character && board[1][2] == character {
		win = true
	}

	if board[2][0] == character && board[2][1] == character && board[2][2] == character {
		win = true
	}

	// Check the rows
	if board[0][0] == character && board[1][0] == character && board[2][0] == character {
		win = true
	}

	if board[0][1] == character && board[1][1] == character && board[2][1] == character {
		win = true
	}

	if board[0][2] == character && board[1][2] == character && board[2][2] == character {
		win = true
	}

	// Check the diagonals
	if board[0][0] == character && board[1][1] == character && board[2][2] == character {
		win = true
	}

	if board[0][2] == character && board[1][1] == character && board[2][0] == character {
		win = true
	}

	if win {
		fmt.Println(fmt.Sprintf("The winning player is player %d", currentPlayer.number))
		os.Exit(1)
	}
}