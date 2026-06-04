package math

import (
	"math"

	"github.com/Matovv/mw4server/internal/pkg/types"
)

func Move(
	currentPos types.Vec2,
	targetPos types.Vec2,
	speed uint64,
	tickRate uint8,
) (types.Vec2, bool) {
	dx := targetPos.X - currentPos.X
	dy := targetPos.Y - currentPos.Y
	if dx == 0 && dy == 0 {
		return targetPos, false
	}
	distance := int64(
		math.Sqrt(
			float64(dx*dx + dy*dy),
		),
	)
	if distance == 0 {
		return targetPos, false
	}
	step := int64(speed) / int64(tickRate)
	// Destination reachable this tick.
	if distance <= step {
		return targetPos, false
	}
	moveX := dx * step / distance
	moveY := dy * step / distance
	if moveX == 0 {
		moveX = sign(dx)
	}
	if moveY == 0 {
		moveY = sign(dy)
	}
	return types.Vec2{
		X: currentPos.X + moveX,
		Y: currentPos.Y + moveY,
	}, true
}

func sign(v int64) int64 {
	switch {
	case v > 0:
		return 1
	case v < 0:
		return -1
	default:
		return 0
	}
}