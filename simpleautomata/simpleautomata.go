package simpleautomata

import (
	"github.com/jemoore/go-cellular-automata/grid"
)

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
	value = value % 256
	// convert value to list of binary digits
	binary_digits := make([]bool, 8)
	for i := 0; i < 8; i++ {
		binary_digits[i] = (value & (1 << uint(i))) != 0
	}
	return binary_digits
}

type SimpleAutomataGame struct {
	grid      grid.Grid
	RuleSet int
}

func (g *SimpleAutomataGame) SetInitialState() {
	g.grid.Cells[grid.COLS/2][0] = true
}

func (g *SimpleAutomataGame) Draw() {
	g.grid.Draw()
}

func (g *SimpleAutomataGame) Update() {
	results := convert_value_to_binary(g.RuleSet)

	for j := int32(1); j < grid.ROWS; j++ {
		for i := int32(0); i < grid.COLS; i++ {
			first := i-1
			if first < 0 {
				first = grid.COLS - 1
			}
			v1 := g.grid.Cells[first%grid.COLS][(j-1)%grid.ROWS]
			v2 := g.grid.Cells[i%grid.COLS][(j-1)%grid.ROWS]
			v3 := g.grid.Cells[(i+1)%grid.COLS][(j-1)%grid.ROWS]
			num := bools_to_number(v1, v2, v3)
			g.grid.Cells[i][j] = results[num]
		}
	}
}

