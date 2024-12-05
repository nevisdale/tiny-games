package state

import (
	"errors"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game interface {
	Update(ctx *StateContext) error
	Draw(ctx *StateContext, screen *ebiten.Image)
}

type StateMode string

const (
	StartMode StateMode = "start"
	GameMode  StateMode = "game"
)

type StateManager map[StateMode]Game

type StateContext struct {
	currentMode   StateMode
	manager       StateManager
	width, height int
}

func NewStateContext(width, height int) StateContext {
	return StateContext{
		manager: make(StateManager),
		width:   width,
		height:  height,
	}
}

func (ctx *StateContext) Register(mode StateMode, state Game) {
	ctx.manager[mode] = state
}

func (ctx *StateContext) SetCurrentMode(mode StateMode) {
	ctx.currentMode = mode
}

func (ctx *StateContext) Update() error {
	g, ok := ctx.manager[ctx.currentMode]
	if !ok {
		return errors.New("not found")
	}
	return g.Update(ctx)
}

func (ctx *StateContext) Draw(screen *ebiten.Image) {
	g, ok := ctx.manager[ctx.currentMode]
	if !ok {
		return
	}
	g.Draw(ctx, screen)
}

func (ctx *StateContext) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return ctx.width, ctx.height
}
