package main

import (
	"./screen"
)

//github.com/tncardoso/gocurses
// http://www.tldp.org/HOWTO/NCURSES-Programming-HOWTO/

func main() {
	screen.Create()
	// Game loop
	var input int
	for {
		if input == 113 /* q - for quit */ {
			break
		}
		screen.Swap()

		input = screen.Listen()
	}

	screen.End()
}
