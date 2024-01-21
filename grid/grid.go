package grid

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CELL_SIZE int32 = 10
const WINDOW_WIDTH int32 = 1200
const WINDOW_HEIGHT int32 = 600
const COLS int32 = WINDOW_WIDTH / CELL_SIZE
const ROWS int32 = WINDOW_HEIGHT / CELL_SIZE


type Grid struct {
	Cells [COLS][ROWS]bool
	Next  [COLS][ROWS]bool
}

// How can we move rendering from here (remove raylib)?
// Have another interface GridRenderer?
// Have this function take an instance of GridRenderer as a parameter?
func (g *Grid) Draw() {
	for i := int32(0); i < COLS; i++ {
		for j := int32(0); j < ROWS; j++ {
			if g.Cells[i][j] {
				rl.DrawRectangle(i*CELL_SIZE, j*CELL_SIZE, CELL_SIZE, CELL_SIZE, rl.White)
			}
		}
	}
}