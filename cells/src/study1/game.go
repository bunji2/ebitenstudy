package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	numUnits int
)

func init() {
	numUnits = screenWidth / unitCellSize * screenHeight / unitCellSize
}

func NewGame() (g *Game) {
	g = &Game{cells: NewCells(numUnits)}
	return
}

type Game struct {
	ctr   int
	cells []Cell
}

func (g *Game) Update() (e error) {
	g.ctr++
	if g.ctr < 0 {
		g.ctr = 0
	}

	pos := -1
	x, y := ebiten.CursorPosition()
	if x >= 0 && x < screenWidth && y >= 0 && y < screenHeight {
		//if x < 0 || x >= screenWidth || y < 0 || y >= screenHeight {
		i := x / unitCellSize
		j := y / unitCellSize
		pos = i + j*screenWidth/unitCellSize
	}

	//if g.ctr%5 == 0 {
	for i := 0; i < numUnits; i++ {
		if i == pos {
			g.cells[i].value = maxValue
		} else {
			g.cells[i].Update()
		}
	}
	//}

	return
}

func (g *Game) Draw(screen *ebiten.Image) {

	numUnits := screenWidth / unitCellSize * screenHeight / unitCellSize
	for i := 0; i < numUnits; i++ {
		g.cells[i].Draw(screen)
	}
	x, y := ebiten.CursorPosition()
	ebitenutil.DebugPrint(screen, fmt.Sprintf("%d,%d", x, y))
	//fmt.Println("[debug]", "nextX =", nextX, "nextY =", nextY) // for debug
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
