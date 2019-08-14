package main

import (
	"fmt"
)

type Board struct {
	Layout [3][3]Tile
}

func NewBoard() Board {
	var b Board
	return b
}

func (b Board) OsWins() bool {
	oWins := false

	if b.oRows() || b.oDiagonals() {
		oWins = true
	}

	return oWins
}

func (b Board) oRows() bool {
	oRows := false

	// Top row
	if b.Layout[0][0].GetState() == "o" && b.Layout[0][1].GetState() == "o" && b.Layout[0][2].GetState() == "o" {
		oRows = true
		// Middle Row
	} else if b.Layout[1][0].GetState() == "o" && b.Layout[1][1].GetState() == "o" && b.Layout[1][2].GetState() == "o" {
		oRows = true
		// Bottom Row
	} else if b.Layout[2][0].GetState() == "o" && b.Layout[2][1].GetState() == "o" && b.Layout[2][2].GetState() == "o" {
		oRows = true
		// Left Column
	} else if b.Layout[0][0].GetState() == "o" && b.Layout[1][0].GetState() == "o" && b.Layout[2][0].GetState() == "o" {
		oRows = true
		// Middle Column
	} else if b.Layout[0][1].GetState() == "o" && b.Layout[1][1].GetState() == "o" && b.Layout[2][1].GetState() == "o" {
		oRows = true
		// Right Column
	} else if b.Layout[0][2].GetState() == "o" && b.Layout[1][2].GetState() == "o" && b.Layout[2][2].GetState() == "o" {
		oRows = true
	}

	return oRows
}

func (b Board) oDiagonals() bool {
	oDiagonals := false

	if b.Layout[0][0].GetState() == "o" && b.Layout[1][1].GetState() == "o" && b.Layout[2][2].GetState() == "o" {
		oDiagonals = true
	} else if b.Layout[0][2].GetState() == "o" && b.Layout[1][1].GetState() == "o" && b.Layout[0][2].GetState() == "o" {
		oDiagonals = true
	}

	return oDiagonals
}

func (b Board) XsWins() bool {
	xWins := false

	if b.xRows() || b.xDiagonals() {
		xWins = true
	}

	return xWins
}

func (b Board) xRows() bool {
	xRows := false

	// Top row
	if b.Layout[0][0].GetState() == "x" && b.Layout[0][1].GetState() == "x" && b.Layout[0][2].GetState() == "x" {
		xRows = true
		// Middle Row
	} else if b.Layout[1][0].GetState() == "x" && b.Layout[1][1].GetState() == "x" && b.Layout[1][2].GetState() == "x" {
		xRows = true
		// Bottom Row
	} else if b.Layout[2][0].GetState() == "x" && b.Layout[2][1].GetState() == "x" && b.Layout[2][2].GetState() == "x" {
		xRows = true
		// Left Column
	} else if b.Layout[0][0].GetState() == "x" && b.Layout[1][0].GetState() == "x" && b.Layout[2][0].GetState() == "x" {
		xRows = true
		// Middle Column
	} else if b.Layout[0][1].GetState() == "x" && b.Layout[1][1].GetState() == "x" && b.Layout[2][1].GetState() == "x" {
		xRows = true
		// Right Column
	} else if b.Layout[0][2].GetState() == "x" && b.Layout[1][2].GetState() == "x" && b.Layout[2][2].GetState() == "x" {
		xRows = true
	}

	return xRows
}

func (b Board) xDiagonals() bool {
	xDiagonals := false

	if b.Layout[0][0].GetState() == "x" && b.Layout[1][1].GetState() == "x" && b.Layout[2][2].GetState() == "x" {
		xDiagonals = true
	} else if b.Layout[0][2].GetState() == "x" && b.Layout[1][1].GetState() == "x" && b.Layout[0][2].GetState() == "x" {
		xDiagonals = true
	}

	return xDiagonals
}

func (b Board) Draw() {
	for row := 0; row < 3; row++ {
		fmt.Printf("+---+---+---+ \n")
		for col := 0; col < 3; col++ {
			fmt.Printf("| %s ", b.Layout[row][col].GetState())
		}
		fmt.Printf("|")
		fmt.Printf("\n")
	}
	fmt.Printf("+---+---+---+ \n")
}
