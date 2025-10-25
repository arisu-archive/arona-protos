package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EquipmentDB struct {
	ConsumableItemBaseDB
	Type                   flatdata.ParcelType
	Level                  int32
	Exp                    int64
	Tier                   int32
	BoundCharacterServerId int64
}
