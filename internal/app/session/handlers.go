package session

import (
    "github.com/Matovv/mw4server/internal/app/command"
)

func (s *Session) handleMove(cmd *command.MoveCommand) {
    move := cmd
    unit := s.World.Units[move.UnitID]
    if unit == nil {
        return
    }
    unit.MoveTarget = move.Position
    unit.Moving = true    
}

func (s *Session) handleAttack(cmd *command.AttackCommand) {
    attack := cmd
    attacker := s.World.Units[attack.UnitID]
    target := s.World.Units[attack.TargetID.ToUnitID()]
    if attacker == nil || target == nil {
        return
    }
    attacker.TargetID = attack.TargetID
}