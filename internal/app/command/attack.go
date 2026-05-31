package command

import "github.com/Matovv/mw4server/internal/pkg/types"

type AttackCommand struct {
	Tick uint64
	PlayerID types.PlayerID
	UnitID   types.UnitID
	TargetID types.TargetID
}

func (*AttackCommand) CommandType() Type {
    return AttackCommandType
}