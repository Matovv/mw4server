package world

import "github.com/Matovv/mw4server/internal/pkg/types"

type Unit struct {
	PrefabID 	  types.PrefabID
	ID            types.UnitID
	OwnerPlayerID types.PlayerID
	Faction       types.Faction
	Position      types.Vec2	
	TargetID      types.TargetID
	Dead          bool
	Hp            uint64
	Mana          uint64	
	Stats         UnitStats
}

type UnitStats struct {
	MaxHp      	uint64
	MaxMana    	uint64
	MoveSpeed  	uint64
}