package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const CELL_SIZE int32 = 10
const WINDOW_WIDTH int32 = 800
const WINDOW_HEIGHT int32 = 800
const COLS int32 = WINDOW_WIDTH / CELL_SIZE
const ROWS int32 = WINDOW_HEIGHT / CELL_SIZE

type GameInterface interface {
	draw()
	update()
	setInitialState()
}

type CellularAutomataGame struct {
	Cells [COLS][ROWS]bool
	Next  [COLS][ROWS]bool
}

func (g *CellularAutomataGame) draw() {
	for i := int32(0); i < COLS; i++ {
		for j := int32(0); j < ROWS; j++ {
			if g.Cells[i][j] {
				rl.DrawRectangle(i*CELL_SIZE, j*CELL_SIZE, CELL_SIZE, CELL_SIZE, rl.White)
			}
		}
	}
}

type GameOfLife struct {
	game CellularAutomataGame
}

func (g *GameOfLife) draw() {
	g.game.draw()
}

func (g *GameOfLife) update() {
	g.game.Next = g.game.Cells

	for i := int32(0); i < COLS; i++ {
		for j := int32(0); j < ROWS; j++ {
			aliveNeighbours := 0
			for x := int32(-1); x <= 1; x++ {
				for y := int32(-1); y <= 1; y++ {
					if !(x == 0 && y == 0) {
						if g.game.Cells[(i+x+COLS)%COLS][(j+y+ROWS)%ROWS] {
							aliveNeighbours++
						}
					}
				}
			}
			if g.game.Cells[i][j] && (aliveNeighbours < 2 || aliveNeighbours > 3) {
				g.game.Next[i][j] = false
			} else if !g.game.Cells[i][j] && aliveNeighbours == 3 {
				g.game.Next[i][j] = true
			}
		}
	}
	g.game.Cells = g.game.Next
}

func (g *GameOfLife) setInitialState() {
	for i := int32(0); i < COLS; i++ {
		for j := int32(0); j < ROWS; j++ {
			g.game.Cells[i][j] = rl.GetRandomValue(0, 1) == 1
		}
	}
}

func bools_to_number(a, b, c bool) int {
	A, B, C := 0, 0, 0
	if a {
		A = 1
	}
	if b {
		B = 1
	}
	if c {
		C = 1
	}
	return 4*A + 2*B + C
}

func convert_value_to_binary(value int) []bool {
	// convert value to list of binary digits
	binary_digits := make([]bool, 8)
	for i := 0; i < 8; i++ {
		binary_digits[i] = (value & (1 << uint(i))) != 0
	}
	return binary_digits
}

type SimpleAutomataGame struct {
	game      CellularAutomataGame
	setNumber int
}

func (g *SimpleAutomataGame) setInitialState() {
	g.game.Cells[COLS/2][0] = true
	g.setNumber = 1
}

func (g *SimpleAutomataGame) draw() {
	g.game.draw()
}

func (g *SimpleAutomataGame) update() {
	results := convert_value_to_binary(g.setNumber)

	for j := int32(1); j < ROWS; j++ {
		for i := int32(0); i < COLS; i++ {
			first := i-1
			if first < 0 {
				first = COLS - 1
			}
			v1 := g.game.Cells[first%COLS][(j-1)%ROWS]
			v2 := g.game.Cells[i%COLS][(j-1)%ROWS]
			v3 := g.game.Cells[(i+1)%COLS][(j-1)%ROWS]
			num := bools_to_number(v1, v2, v3)
			g.game.Cells[i][j] = results[num]
		}
	}
}

func drawGrid() {
	for i := int32(0); i < COLS; i++ {
		rl.DrawLine(i*CELL_SIZE, 0, i*CELL_SIZE, WINDOW_HEIGHT, rl.LightGray)
		rl.DrawLine(0, i*CELL_SIZE, WINDOW_WIDTH, i*CELL_SIZE, rl.LightGray)
	}
}

func main() {
	rl.InitWindow(WINDOW_WIDTH, WINDOW_HEIGHT, "Game of Life")
	rl.SetTargetFPS(10)

	// var game GameInterface = &GameOfLife{}
	var game GameInterface = &SimpleAutomataGame{}
	game.setInitialState()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		game.draw()
		drawGrid()
		game.update()
		rl.EndDrawing()
	}
}
