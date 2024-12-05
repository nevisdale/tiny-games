package vec2

import (
	"math"
)

type Vec2 struct {
	X, Y float64
}

func (v Vec2) Add(vec Vec2) Vec2 {
	v.X += vec.X
	v.Y += vec.Y
	return v
}

func (v Vec2) Sub(vec Vec2) Vec2 {
	v.X -= vec.X
	v.Y -= vec.Y
	return v
}

func (v Vec2) Scale(n float64) Vec2 {
	v.X *= n
	v.Y *= n
	return v
}

func (v Vec2) MagSq() float64 {
	return v.X*v.X + v.Y*v.Y
}

func (v Vec2) Mag() float64 {
	return math.Sqrt(v.MagSq())
}

func (v Vec2) Equal(vec Vec2) bool {
	return v.X == vec.X && v.Y == vec.Y
}

func (v Vec2) IsZero() bool {
	return v.Equal(Vec2{})
}
