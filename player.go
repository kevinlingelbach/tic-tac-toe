package main

import (
	"fmt"
)

type Player struct {
	Turns int
}

func NewPlayer() Player {
	var p Player
	p.Turns = 0
	return p
}

func (p *Player) Turn(b *Board) {
	if p.Turns%2 == 0 {
		fmt.Printf("X's turn! \n")
	} else {
		fmt.Printf("O's turn! \n")
	}

	properInput := false

	for !properInput {
		row, col := getPosition()

		if b.Layout[row][col].GetState() == "x" || b.Layout[row][col].GetState() == "o" {
			fmt.Printf("Please pick a position with no x or o. \n")
		} else {
			if p.Turns%2 == 0 {
				b.Layout[row][col].ChangeToX()
			} else {
				b.Layout[row][col].ChangeToO()
			}
			properInput = true
		}
	}

	p.Turns++
}

func getPosition() (int, int) {
	var row, col int
	properInput := false

	for !properInput {
		fmt.Printf("Select the position you want to change...\n")
		fmt.Printf("Enter the row: ")
		fmt.Scan(&row)
		fmt.Printf("Enter the column: ")
		fmt.Scan(&col)

		if row < 0 || row > 2 || col < 0 || col > 2 {
			fmt.Printf("Please enter only numbers from 0 to 2. \n")
		} else {
			properInput = true
		}
	}

	return row, col
}
