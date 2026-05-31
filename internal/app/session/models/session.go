package models

import (
	"time"

	"github.com/Matovv/mw4server/internal/pkg/types"
)

type PlayerSlot struct {
	Index      uint8
	PlayerID   types.PlayerID
	Connected  bool
	LastSeen   time.Time
	HeroUnitID types.UnitID
}