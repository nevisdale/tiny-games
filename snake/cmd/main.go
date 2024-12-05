package main

import (
	"snake/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	width  = 640
	height = 640
)

func main() {
	ebiten.SetWindowTitle("snake for Alina")
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)

	game := game.New(width, height)
	if err := ebiten.RunGame(&game); err != nil {
		panic(err)
	}
}
