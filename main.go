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
		//exit, _ := strconv.Atoi("q")
		gocurses.Addch('q')
		if input == 113 /* Q */ {
			break
		}
		screen.Write("Match 3!", 20, 1)
		screen.Swap()
		wind := screen.Window{}
		wind.Height = 10
		wind.Width = 10
		wind.X = 10
		wind.Y = 10
		wind.Make()
		wind.Draw()
		input = screen.Listen()
	}

	screen.End()

}
