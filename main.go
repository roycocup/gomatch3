package main

import (
	"strconv"

	"./debug"
	"./screen"
)

//github.com/tncardoso/gocurses

func main() {
	screen.Create()
	// Game loop
	var input int
	for {
		if input == 113 /* q - for quit */ {
			break
		}
		screen.Swap()

		d := debug.Debug{}
		debug.AddDebug(&d, "add my debug ")
		debug.AddDebug(&d, "col - "+strconv.Itoa(screen.MaxCols()))
		debug.AddDebug(&d, "row - "+strconv.Itoa(screen.MaxRows()))
		input = screen.Listen()
	}

	screen.End()
}
