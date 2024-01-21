package main

import (
	"flag"

	rl "github.com/gen2brain/raylib-go/raylib"
	gol "github.com/jemoore/go-cellular-automata/gameoflife"
	"github.com/jemoore/go-cellular-automata/grid"
	"github.com/jemoore/go-cellular-automata/interfaces"
	sa "github.com/jemoore/go-cellular-automata/simpleautomata"
)

func drawGrid() {
	for i := int32(0); i < grid.COLS; i++ {
		rl.DrawLine(i*grid.CELL_SIZE, 0, i*grid.CELL_SIZE, grid.WINDOW_HEIGHT, rl.LightGray)
		rl.DrawLine(0, i*grid.CELL_SIZE, grid.WINDOW_WIDTH, i*grid.CELL_SIZE, rl.LightGray)
	}
}

func main() {
	simple := flag.Bool("simple", false, "Use simple automata")
	ruleSet := flag.Int("rule", 110, "Rule set to use")
	flag.Parse()

	rl.InitWindow(grid.WINDOW_WIDTH, grid.WINDOW_HEIGHT, "Game of Life")
	rl.SetTargetFPS(10)

	var game interfaces.GameInterface
	if *simple {
		game = &sa.SimpleAutomataGame{ RuleSet : *ruleSet }
	} else {
		game = &gol.GameOfLife{}
	}

	game.SetInitialState()

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		game.Draw()
		drawGrid()
		game.Update()
		rl.EndDrawing()
	}
}
