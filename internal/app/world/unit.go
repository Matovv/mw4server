package world

import "github.com/Matovv/mw4server/internal/pkg/types"

type Unit struct {
	ID            types.UnitID
	OwnerPlayerID types.PlayerID
	Faction       types.Faction
	Position      types.Vec2	
	TargetID      types.TargetID
	Dead          bool
	HP            int64
	Mana          int64	
	Stats         UnitStats
}

type UnitStats struct {
	MaxHP      	int64
	MaxMana    	int64
	MoveSpeed  	int64
}