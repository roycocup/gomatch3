package screen

import "github.com/tncardoso/gocurses"

type Window struct {
	Width  int
	Height int
	X      int
	Y      int
	Wind   gocurses.Window
}

func (window *Window) Draw() {
	window.Wind.Refresh()
}

func (window *Window) Make() {
	wind := gocurses.NewWindow(window.Height, window.Width, window.Y, window.X)
	wind.Box(0, 0)
}
