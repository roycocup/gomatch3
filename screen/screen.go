package screen

import "github.com/tncardoso/gocurses"

func Create() {
	gocurses.Initscr()
	gocurses.Cbreak()
	gocurses.Noecho()
	gocurses.Stdscr.Keypad(true)
}

func Swap() {
	gocurses.Refresh()
}

func Listen() int {
	return gocurses.Stdscr.Getch()
}

func End() {
	gocurses.End()
}

func (window *Window) NewWindow(w int, h int, x int, y int) *Window {
	return new(Window)
}
