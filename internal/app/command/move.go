package command

import "github.com/Matovv/mw4server/internal/pkg/types"

type MoveCommand struct {
	Tick uint64
	PlayerID types.PlayerID
	UnitID types.UnitID
	Position types.Vec2
}

func (*MoveCommand) CommandType() Type {
    return MoveCommandType
}