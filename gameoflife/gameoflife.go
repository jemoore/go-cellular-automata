package gameoflife

import (
	"math/rand"
	"time"

	"github.com/jemoore/go-cellular-automata/grid"
)

func init() {
}

type GameOfLife struct {
	grid grid.Grid
}

func (g *GameOfLife) SetInitialState() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	for i := int32(0); i < grid.COLS; i++ {
		for j := int32(0); j < grid.ROWS; j++ {
			g.grid.Cells[i][j] = r1.Intn(2) == 1
		}
	}
}

func (g *GameOfLife) Draw() {
	g.grid.Draw()
}

func (g *GameOfLife) Update() {
	g.grid.Next = g.grid.Cells

	for i := int32(0); i < grid.COLS; i++ {
		for j := int32(0); j < grid.ROWS; j++ {
			aliveNeighbours := 0
			for x := int32(-1); x <= 1; x++ {
				for y := int32(-1); y <= 1; y++ {
					if !(x == 0 && y == 0) {
						if g.grid.Cells[(i+x+grid.COLS)%grid.COLS][(j+y+grid.ROWS)%grid.ROWS] {
							aliveNeighbours++
						}
					}
				}
			}
			if g.grid.Cells[i][j] && (aliveNeighbours < 2 || aliveNeighbours > 3) {
				g.grid.Next[i][j] = false
			} else if !g.grid.Cells[i][j] && aliveNeighbours == 3 {
				g.grid.Next[i][j] = true
			}
		}
	}
	g.grid.Cells = g.grid.Next
}
