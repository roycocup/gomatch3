package debug

import (
	"fmt"

	"github.com/tncardoso/gocurses"
)

func PrintV(o interface{}) {
	fmt.Print("\n\n")
	fmt.Print("####################\n")
	fmt.Printf("%#v", o)
	fmt.Print("####################\n")
	fmt.Print("\n")
}

type Debug struct {
	x       int
	y       int
	h       int
	w       int
	entries int
	window  gocurses.Window
	strings []string
}

func AddDebug(debug *Debug, str string) {
	debug.strings = append(debug.strings, str)
	debug.entries++
	debug.h = debug.entries + 2
	debug.w = len(str) + 2

	debug.window = *gocurses.NewWindow(debug.h, debug.w, 0, 0)
	debug.window.Box(1, 20)

	i := 1
	for _, v := range debug.strings {
		strY := debug.y + i
		strX := debug.x + 1
		debug.window.Mvaddstr(strY, strX, v)
		i++
	}

	debug.window.Refresh()
}
