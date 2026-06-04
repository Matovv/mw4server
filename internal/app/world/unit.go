package world

import (
	"github.com/Matovv/mw4server/internal/app/protocol"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type Unit struct {
	GameModelID   string
	Size          uint64
	ID            types.UnitID
	OwnerPlayerID types.PlayerID
	Faction       types.Faction
	Position      types.Vec2	
	TargetID      types.TargetID
    MoveTarget 	  types.Vec2
	Moving        bool
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

func (u *Unit) ToSnapshot() protocol.UnitSnapshot {
	return protocol.UnitSnapshot{
		GameModelID: u.GameModelID,
		Size: u.Size,
		ID: uint64(u.ID),
		Faction: string(u.Faction),
		X: u.Position.X,
		Y: u.Position.Y,
		Hp: u.Hp,
		Mana: u.Mana,
    }
}