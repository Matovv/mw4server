package world

import (
	"github.com/Matovv/mw4server/internal/app/command"
)

type NeutralAI struct {
    PlayerID uint64
}


func (ai *NeutralAI) Think(w *World) []command.Command {
    return nil
}