package screen

import "github.com/tncardoso/gocurses"

type Window struct {
	Width  int
	Height int
	X      int
	Y      int
	wind   gocurses.Window
}

func (window *Window) Draw() {
	window.wind.Refresh()
}

func (window *Window) Make() {
	window.wind = *gocurses.NewWindow(window.Height, window.Width, window.Y, window.X)
	window.wind.Box(0, 0)
}
