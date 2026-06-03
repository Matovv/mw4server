package world

import (
	"github.com/Matovv/mw4server/internal/pkg/errors"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type World struct {
	MapName       string
	PrefabManager prefabManager
	Units   	  map[types.UnitID]*Unit
	Objects       map[types.ObjectID]*Object
	NeutralAI     *NeutralAI
	NextUnitID    types.UnitID
}

func NewWorld(mapName string, prefabManager prefabManager) *World {
	return &World{
		MapName:   mapName,
		PrefabManager: prefabManager,
		Units:     make(map[types.UnitID]*Unit, 256),
		Objects:   make(map[types.ObjectID]*Object, 256),
		NeutralAI: &NeutralAI{},
		NextUnitID: 1,
	}
}

func (w *World) SpawnUnit(
	prefabID types.PrefabID,
	faction types.Faction,
	position types.Vec2,
) (*Unit, error) {
	prefab := w.PrefabManager.GetPrefabUnit(
		prefabID,
	)
	if prefab == nil {
		return nil, errors.ErrPrefabUnitNotFound
	}
	unitID := w.NewUnitID()
	unit := &Unit{
		ID:            	unitID,
		PrefabID:      	prefabID,
		Faction: 	    faction,
		Position:      	position,
		Hp:   			prefab.MaxHp,
		Mana: 			prefab.MaxMana,
		Stats: UnitStats{
			MaxHp:     	prefab.MaxHp,
			MaxMana:   	prefab.MaxMana,
			MoveSpeed: 	prefab.MoveSpeed,
		},
	}
	w.Units[unitID] = unit
	return unit, nil
}
func (w *World) NewUnitID() types.UnitID {
	id := w.NextUnitID
	w.NextUnitID++
	return id
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