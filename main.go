package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	playing := true

	for playing {
		clear()

		p := Player{}
		b := Board{}

		b.Draw()

		for !b.XsWins() && !b.OsWins() {
			p.Turn(&b)
			clear()
			b.Draw()
		}

		if b.OsWins() {
			fmt.Printf("The Os win! \n")
		} else if b.XsWins() {
			fmt.Printf("The Xs win! \n")
		}

		playing = playAgain()
	}
}

func playAgain() bool {
	playAgain := false
	var response string

	fmt.Printf("Do you want to play again? (y/n)")
	fmt.Scan(&response)

	if response == "y" {
		playAgain = true
	}

	return playAgain
}

func clear() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}
