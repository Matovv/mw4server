package world

import "github.com/Matovv/mw4server/internal/pkg/types"

type Object struct {
	ID uint64

	Type ObjectType

	Position types.Vec2
}

type ObjectType uint8

const (
	Tower ObjectType = iota
	Tree
	Barracks
	Fountain
)