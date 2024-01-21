package grid


const CELL_SIZE int32 = 10
const WINDOW_WIDTH int32 = 800
const WINDOW_HEIGHT int32 = 800
const COLS int32 = WINDOW_WIDTH / CELL_SIZE
const ROWS int32 = WINDOW_HEIGHT / CELL_SIZE


type CellularAutomataGame struct {
	Cells [COLS][ROWS]bool
	Next  [COLS][ROWS]bool
}

func (g *CellularAutomataGame) draw() {
	for i := int32(0); i < COLS; i++ {
		for j := int32(0); j < ROWS; j++ {
			if g.Cells[i][j] {
				// rl.DrawRectangle(i*CELL_SIZE, j*CELL_SIZE, CELL_SIZE, CELL_SIZE, rl.White)
			}
		}
	}
}