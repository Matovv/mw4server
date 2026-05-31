package models

import "github.com/Matovv/mw4server/internal/pkg/types"

type PlayerSlot struct {
	Index      uint8
	PlayerID   types.PlayerID
	Connected  bool
	HeroUnitID types.UnitID
}