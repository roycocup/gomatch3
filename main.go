package main

import (
	"./screen"
	"github.com/tncardoso/gocurses"
)

//github.com/tncardoso/gocurses

func main() {
	screen.Create()

	// Game loop
	var input int
	for {
		if input == 113 /* Q */ {
			break
		}
		screen.Write("Match 3!", 20, 1)
		screen.Swap()
		wind := screen.Window{100, 100, 10, 10, gocurses.Window{}}
		wind.Make()
		wind.Draw()
		input = screen.Listen()
	}

	screen.End()

}
