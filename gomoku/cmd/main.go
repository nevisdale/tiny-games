package main

import (
	"gomoku/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	game := game.NewGame()
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
