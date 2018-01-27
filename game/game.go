package game

import (
	"math/rand"
)

type Game struct {
	grid Grid
}

type Grid struct {
	cols  int
	rows  int
	cells []Cell
}

type Cell struct {
	cellType string
	row      int
	col      int
}

func NewGame() Game {
	grid := getGrid()
	return Game{grid}
}

func getGrid() Grid {
	grid := Grid{}

	grid.cols = 10
	grid.rows = 10

	for col := 0; col < grid.cols; col++ {
		for row := 0; row < grid.rows; row++ {
			cell := getRandomCell()
			cell.row = row
			cell.col = col
			grid.cells = append(grid.cells, cell)
		}
	}

	return grid
}

func getRandomCell() Cell {
	var cells []Cell
	cells = append(cells, Cell{"blister", 0, 0})
	cells = append(cells, Cell{"normal", 0, 0})
	cells = append(cells, Cell{"red", 0, 0})
	cells = append(cells, Cell{"blue", 0, 0})
	cells = append(cells, Cell{"yellow", 0, 0})

	return cells[rand.Intn(len(cells))]
}

func (game Game) Update() {

}

func (game Game) Draw() {

}
