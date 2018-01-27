package screen

import "github.com/tncardoso/gocurses"

func Write(text string, x int, y int) {
	gocurses.Attron(gocurses.A_BOLD)
	gocurses.Mvaddstr(y, x, text)
}
