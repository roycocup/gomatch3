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

func NewWindow(h int, w int, x int, y int) Window {
	wind := Window{}
	wind.Height = h
	wind.Width = w
	wind.X = x
	wind.Y = y
	wind.Make()
	return wind
}

func MaxCols() int {
	cols, _ := gocurses.Getmaxyx()
	return cols
}

func MaxRows() int {
	_, rows := gocurses.Getmaxyx()
	return rows
}
