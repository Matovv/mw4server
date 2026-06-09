package math

import (
	"math"

	"github.com/Matovv/mw4server/internal/pkg/types"
)

func Move(
	currentT types.Transform,
	targetPos types.Vec2,
	speed uint64,
	tickRate uint8,
) (types.Transform, bool) {
	dx := targetPos.X - currentT.Position.X
	dy := targetPos.Y - currentT.Position.Y
	if dx == 0 && dy == 0 {
		return currentT, false
	}
	distance := int64(
		math.Sqrt(
			float64(dx*dx + dy*dy),
		),
	)
	if distance == 0 {
		return currentT, false
	}
	rotation := calculateRotation(
		dx,
		dy,
	)
	currentT.Rotation = rotation
	step := int64(speed) / int64(tickRate)
	if distance <= step {
		currentT.Position = targetPos
		return currentT, false
	}
	moveX := dx * step / distance
	moveY := dy * step / distance
	if moveX == 0 && dx != 0 {
		moveX = sign(dx)
	}
	if moveY == 0 && dy != 0 {
		moveY = sign(dy)
	}
	currentT.Position.X += moveX
	currentT.Position.Y += moveY
	return currentT, true
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

// Rotation cycle - 0: right, 90: down, 180: right, 270: up
func calculateRotation(
    dx int64,
    dy int64,
) types.Rotation {
    angle := math.Atan2(
        float64(-dy),
        float64(dx),
    )
    degrees := angle * 180 / math.Pi
    if degrees < 0 {
        degrees += 360
    }
    return types.Rotation(
        uint16(degrees),
    )
}