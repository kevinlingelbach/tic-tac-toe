package main

type Tile struct {
	state string
}

func NewTile() Tile {
	var t Tile
	t.state = "0"
	return t
}

func (t *Tile) ChangeToX() {
	t.state = "x"
}

func (t *Tile) ChangeToO() {
	t.state = "o"
}

func (t Tile) DrawTile() {
	if t.state == "0" {

	} else if t.state == "o" {

	} else if t.state == "x" {

	}
}

func (t Tile) GetState() string {
	state := " "
	if t.state == "x" || t.state == "o" {
		state = t.state
	}

	return state
}
