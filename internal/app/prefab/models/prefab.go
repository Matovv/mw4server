package models

import "github.com/Matovv/mw4server/internal/pkg/types"

type PrefabUnit struct {
	PrefabID    types.PrefabID
	GameModelID string
	Size        uint64
	MaxHp       uint64
	MaxMana     uint64
	MoveSpeed   uint64
}