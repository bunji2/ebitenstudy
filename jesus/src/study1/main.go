package main

import (
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	windowTitle  = "Jesus (Study1)"
	screenWidth  = 640
	screenHeight = 480
)

func main() {
	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle(windowTitle)

	err := ebiten.RunGame(NewGame())
	if err != nil {
		log.Fatal(err)
	}
}
