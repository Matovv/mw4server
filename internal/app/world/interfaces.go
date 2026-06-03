package world

import (
	"github.com/Matovv/mw4server/internal/app/prefab/models"
	"github.com/Matovv/mw4server/internal/pkg/types"
)

type prefabManager interface {
	GetPrefabUnit(prefabId types.PrefabID) *models.PrefabUnit
}