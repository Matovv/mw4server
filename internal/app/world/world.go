package world

import "github.com/Matovv/mw4server/internal/pkg/types"

type World struct {
	MapName string
	Units   map[types.UnitID]*Unit
	Objects map[types.ObjectID]*Object
	NeutralAI *NeutralAI
}

func NewWorld(mapName string) *World {
	return &World{
		MapName:   mapName,
		Units:     make(map[types.UnitID]*Unit),
		Objects:   make(map[types.ObjectID]*Object),
		NeutralAI: &NeutralAI{},
	}
}

func (w *World) Update() {
    // later:
	// movement
	// combat
	// cooldowns
	// AI

	if w == nil {
		return
	}
}