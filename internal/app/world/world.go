package world

import "github.com/Matovv/mw4server/internal/pkg/types"

type World struct {
	MapName       string
	PrefabManager prefabManager
	Units   	  map[types.UnitID]*Unit
	Objects       map[types.ObjectID]*Object
	NeutralAI     *NeutralAI
}

func NewWorld(mapName string, prefabManager prefabManager) *World {
	return &World{
		MapName:   mapName,
		PrefabManager: prefabManager,
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