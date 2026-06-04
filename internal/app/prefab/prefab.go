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
		Units: make(map[types.PrefabID]models.PrefabUnit, 64),
	}
}

func (m *PrefabManager) GetPrefabUnit(prefabId types.PrefabID) *models.PrefabUnit {
	unit, ok := m.Units[prefabId]
	if ok {
		return &unit
	}
	return nil
}

func (m *PrefabManager) InitTestData() {
	prefabUnit1 := models.PrefabUnit{
		PrefabID: 1,
		GameModelID: "test1",
		Size: 100,
		MaxHp: 300,
		MaxMana: 100,
		MoveSpeed: 300,
	}
	m.Units[prefabUnit1.PrefabID] = prefabUnit1
}