package game

import (
	"fmt"
	"gomoku/internal/board"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Player struct {
	selectedClr color.Color
	clr         color.Color
	boardCell   uint8
}

type Game struct {
	board   *board.Board
	players []Player

	selectedCellX int
	selectedCellY int

	playerTurn                             int     // index of the player in the players slice
	winner                                 int     // index of the player in the players slice
	winnerX0, winnerY0, winnerX1, winnerY1 int     // winning line coordinates on the board
	winnerLineProgress                     float32 // winning line drawing progress
	winnerLineProgressSpeed                float32 // winning line drawing speed

	cellSize int

	bgColor color.Color
}

func NewGame() *Game {
	players := []Player{
		{clr: color.RGBA{255, 0, 0, 255}, boardCell: 1, selectedClr: color.RGBA{255, 0, 0, 100}},
		{clr: color.RGBA{0, 0, 255, 255}, boardCell: 2, selectedClr: color.RGBA{0, 0, 255, 100}},
	}

	g := &Game{
		board:                   board.NewBoard(),
		selectedCellX:           -1,
		selectedCellY:           -1,
		players:                 players,
		playerTurn:              0,
		winner:                  -1,
		winnerLineProgressSpeed: 0.1,

		cellSize: 16,
		bgColor:  color.RGBA{128, 128, 128, 255},
	}

	return g
}

func (g *Game) Update() error {
	mouseX, mouseY := ebiten.CursorPosition()
	g.selectedCellX, g.selectedCellY = int(mouseX/g.cellSize), int(mouseY/g.cellSize)

	if ebiten.IsKeyPressed(ebiten.KeyR) {
		g.board.Reset()
		g.winner = -1
		g.playerTurn = 0
		g.winnerLineProgress = 0
	}

	if g.winner != -1 {
		return nil
	}

	if cell, ok := g.board.Get(g.selectedCellX, g.selectedCellY); cell == 0 && ok {
		if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
			g.board.Set(g.selectedCellX, g.selectedCellY, g.players[g.playerTurn].boardCell)
			var ok bool
			g.winnerX0, g.winnerY0, g.winnerX1, g.winnerY1, ok = g.board.IsWinnerAfterSet(g.selectedCellX, g.selectedCellY)
			// keep the coordinates in order
			// just to draw the winning line from left to right
			if g.winnerX1 < g.winnerX0 || (g.winnerX1 == g.winnerX0 && g.winnerY1 < g.winnerY0) {
				g.winnerX0, g.winnerX1 = g.winnerX1, g.winnerX0
				g.winnerY0, g.winnerY1 = g.winnerY1, g.winnerY0
			}

			if ok {
				g.winner = g.playerTurn
				return nil
			}
			g.playerTurn++
			g.playerTurn %= len(g.players)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.drawGrid(screen)

	// draw selected cell
	cell, ok := g.board.Get(g.selectedCellX, g.selectedCellY)
	if cell == 0 && ok { // could be selected
		g.drawCell(screen, g.selectedCellX, g.selectedCellY, g.players[g.playerTurn].selectedClr)
	}

	w, h := g.board.Size()
	for i := 0; i < w; i++ {
		for j := 0; j < h; j++ {
			cell, _ := g.board.Get(i, j)
			if cell != 0 {
				g.drawCell(screen, i, j, g.players[cell-1].clr)
			}
		}
	}

	if g.winner != -1 {
		ebitenutil.DebugPrintAt(screen, fmt.Sprintf("Player%d won. Press R to restart", g.winner+1), 0, 0)
		srcX := float32(g.winnerX0*g.cellSize) + float32(g.cellSize)/2
		srcY := float32(g.winnerY0*g.cellSize) + float32(g.cellSize)/2
		dstX := float32(g.winnerX1*g.cellSize) + float32(g.cellSize)/2
		dstY := float32(g.winnerY1*g.cellSize) + float32(g.cellSize)/2
		// interpolate the line drawing
		dstX = srcX + (dstX-srcX)*g.winnerLineProgress
		dstY = srcY + (dstY-srcY)*g.winnerLineProgress

		vector.StrokeLine(screen, srcX, srcY, dstX, dstY, 2, color.Black, false)
		g.winnerLineProgress += g.winnerLineProgressSpeed
		if g.winnerLineProgress > 1 {
			g.winnerLineProgress = 1
		}
	}
}

func (g Game) Layout(_, _ int) (_, _ int) {
	w, h := g.board.Size()
	return w * int(g.cellSize), h * int(g.cellSize)
}

func (g Game) drawGrid(screen *ebiten.Image) {
	screen.Fill(g.bgColor)
	w, h := g.board.Size()
	for i := 0; i <= w; i++ {
		vector.StrokeLine(screen,
			float32(i*g.cellSize), 0,
			float32(i*g.cellSize), float32(h*g.cellSize),
			1, color.Black, false) // vertical lines
	}
	for i := 0; i <= h; i++ {
		vector.StrokeLine(screen,
			0, float32(i*g.cellSize),
			float32(w*g.cellSize), float32(i*g.cellSize),
			1, color.Black, false) // horizontal lines
	}
}

func (g Game) drawCell(screen *ebiten.Image, x, y int, clr color.Color) {
	vector.DrawFilledRect(screen,
		float32(x*g.cellSize), float32(y*g.cellSize),
		float32(g.cellSize)-1, float32(g.cellSize)-1, // minus 1 to avoid overlapping with grid (stroke line width)
		clr, false)
}
