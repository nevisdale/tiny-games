package game

import (
	"fmt"
	"image/color"
	"math/rand"
	"snake/pkg/vec2"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	width, height int

	snakeBody []vec2.Vec2

	velocity vec2.Vec2

	updateHeadTimer *time.Ticker
	speedTimers     map[int]time.Duration

	berryCount   int
	berries      []vec2.Vec2
	moveOnlyHead bool

	ceilSize int

	score  int
	isOver bool
}

func (g *Game) reset() {
	g.snakeBody = []vec2.Vec2{{X: float64(g.width) / 2, Y: float64(g.height) / 2}}
	g.velocity = vec2.Vec2{X: 1, Y: 0}
	g.updateHeadTimer.Reset(100 * time.Millisecond)
	g.score = 0
	g.isOver = false
}

func New(width, height int) Game {
	g := Game{
		// screen settings
		width:  width,
		height: height,

		// snake
		snakeBody:       []vec2.Vec2{{X: float64(width) / 2, Y: float64(height) / 2}},
		velocity:        vec2.Vec2{X: 1, Y: 0},
		updateHeadTimer: time.NewTicker(100 * time.Millisecond),

		berryCount: 1,

		// board settings
		ceilSize: 20,

		speedTimers: map[int]time.Duration{
			0:  100 * time.Millisecond,
			3:  80 * time.Millisecond,
			6:  60 * time.Millisecond,
			9:  50 * time.Millisecond,
			12: 40 * time.Millisecond,
			15: 30 * time.Millisecond,
			18: 20 * time.Millisecond,
			21: 15 * time.Millisecond,
			24: 10 * time.Millisecond,
		},
	}

	return g
}

func (g *Game) Update() error {
	if g.isOver {
		if ebiten.IsKeyPressed(ebiten.KeyP) {
			g.reset()
		}
		return nil
	}

	// get direction to move
	if direction := getInput(); !direction.IsZero() {
		switch {
		// opposite directions
		case g.velocity.X != 0 && g.velocity.X+direction.X == 0:
		case g.velocity.Y != 0 && g.velocity.Y+direction.Y == 0:
		default:
			g.velocity = direction
		}
	}

	// update snake position
	select {
	case <-g.updateHeadTimer.C:
		realVelocity := g.velocity.Scale(float64(g.ceilSize))

		head := g.snakeBody[0]
		newHead := head.Add(realVelocity)

		if newHead.X < 0 || newHead.X >= float64(g.width) || newHead.Y < 0 || newHead.Y >= float64(g.height) {
			g.isOver = true
			break
		}

		if !g.moveOnlyHead {
			for i := len(g.snakeBody) - 1; i > 0; i-- {
				g.snakeBody[i] = g.snakeBody[i-1]
			}
		}

		g.snakeBody[0] = newHead
		g.moveOnlyHead = false

	default:
		if newDuration, ok := g.speedTimers[g.score]; ok {
			g.updateHeadTimer.Reset(newDuration)
			delete(g.speedTimers, g.score)
		}
	}

	// generate berries
	for i := 0; i < g.berryCount-len(g.berries); i++ {
		berryPosition := vec2.Vec2{
			X: float64(rand.Intn(g.width) / g.ceilSize * g.ceilSize),
			Y: float64(rand.Intn(g.height) / g.ceilSize * g.ceilSize),
		}
		g.berries = append(g.berries, berryPosition)
	}

	// berry x snake collision
	{
		head := g.snakeBody[0]
		for i, berry := range g.berries {
			if berry.Equal(head) {
				g.berries = append(g.berries[:i], g.berries[i+1:]...)
				g.score += 1
				g.snakeBody = append([]vec2.Vec2{head}, g.snakeBody...)
				g.moveOnlyHead = true
			}
		}
	}

	// snake collision
	{
		// without head
		head := g.snakeBody[0]
		if len(g.snakeBody) > 3 {
			for _, bodyCell := range g.snakeBody[3:] {
				if bodyCell.Equal(head) {
					g.isOver = true
				}
			}
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{R: 0, B: 100, G: 0})
	ebitenutil.DebugPrint(screen, fmt.Sprintf("score: %d\n", g.score))

	if g.isOver {
		ebitenutil.DebugPrint(screen, "\nGAME OVER")
		ebitenutil.DebugPrint(screen, "\n\npress P to RESTART")
	}

	for _, bodyCell := range g.snakeBody {
		vector.DrawFilledRect(
			screen,
			float32(bodyCell.X),
			float32(bodyCell.Y),
			float32(g.ceilSize),
			float32(g.ceilSize),
			color.White,
			true)
	}

	for _, berry := range g.berries {
		vector.DrawFilledRect(
			screen,
			float32(berry.X),
			float32(berry.Y),
			float32(g.ceilSize),
			float32(g.ceilSize),
			color.RGBA{R: 0, G: 255, B: 0},
			true,
		)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func getInput() vec2.Vec2 {
	direction := vec2.Vec2{}
	switch {
	case inpututil.IsKeyJustPressed(ebiten.KeyA):
		direction.X -= 1
	case inpututil.IsKeyJustPressed(ebiten.KeyD):
		direction.X += 1
	case inpututil.IsKeyJustPressed(ebiten.KeyW):
		direction.Y -= 1
	case inpututil.IsKeyJustPressed(ebiten.KeyS):
		direction.Y += 1
	}
	return direction
}
