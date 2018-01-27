package main

import (
	"./game"
	"./screen"
)

// github.com/tncardoso/gocurses
// http://www.tldp.org/HOWTO/NCURSES-Programming-HOWTO/
type gameItem interface {
	Update()
	Draw()
}

var dependencies []gameItem

func main() {
	screen.Create()

	dependencies = append(dependencies, game.NewGame())

	loop(dependencies)

	screen.End()
}

func loop(dependencies []gameItem) int {
	for {
		update(dependencies)
		draw(dependencies)
		if screen.Listen() == 113 {
			break
		}
	}
	return 0
}

func update(deps []gameItem) {
	for _, v := range deps {
		v.Update()
	}

}

func draw(deps []gameItem) {
	for _, v := range deps {
		v.Draw()
	}
	screen.Swap()
}
