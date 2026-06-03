package prefab

import (
	"github.com/Matovv/mw4server/internal/app/prefab/models"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type PrefabManager struct {
	Units map[types.PrefabID]models.PrefabUnit
}

func NewPrefabManager() *PrefabManager {
	return &PrefabManager{
		Units: make(map[types.PrefabID]models.PrefabUnit, 100),
	}
}

func (m *PrefabManager) GetPrefabUnit(prefabId types.PrefabID) *models.PrefabUnit {
	unit, ok := m.Units[prefabId]
	if ok {
		return &unit
	}
	return nil
}