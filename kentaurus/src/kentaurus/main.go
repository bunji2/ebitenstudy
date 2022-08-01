package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowTitle  = "Kentaurus (Study1)"
	screenWidth  = 780
	screenHeight = 546
	centerX      = screenWidth / 2
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(windowTitle)

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
