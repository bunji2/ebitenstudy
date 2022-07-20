package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const (
	unitCellSize = 40
	maxValue     = 100
)

var (
	activeColor color.RGBA
)

func init() {
	activeColor = color.RGBA{R: 0xFF, A: 0xFF}
}

func NewCells(size int) (cells []Cell) {
	cells = make([]Cell, size)
	for pos := 0; pos < size; pos++ {
		cells[pos] = Cell{
			pos: pos,
		}
	}
	return
}

type Cell struct {
	pos   int //  0 ≦ pos < screenWidth/unitCellSize*screenHeight/unitCellSize
	value int // 0 ≦ value ≦ 100
}

func (c *Cell) XY() (x, y float64) {
	wCount := screenWidth / unitCellSize
	x = float64(unitCellSize) * float64(c.pos%wCount)
	y = float64(unitCellSize) * float64(c.pos/wCount)
	return
}

func (c *Cell) Update() (e error) {
	if c.value == 0 {
		return
	}

	c.value--
	if c.value < 0 {
		c.value = 0
	}
	return
}

func (c *Cell) Draw(screen *ebiten.Image) {
	if c.value == 0 {
		return
	}
	x, y := c.XY()
	n := 0xFF * c.value / maxValue
	clr := color.RGBA{
		R: uint8(n),
		G: uint8(n),
		B: uint8(n),
		A: 0xFF,
	}

	ebitenutil.DrawRect(screen, x, y, unitCellSize-1, unitCellSize-1, clr)

	/*
		上記は次のように記述しても同じ結果となった

		img := ebiten.NewImage(unitCellSize, unitCellSize)
		ebitenutil.DrawRect(img, 0, 0, unitCellSize-1, unitCellSize-1, clr)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(x, y)
		screen.DrawImage(img, op)
	*/
}
