package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type GearDB struct {
	ParcelBase
	Type                   flatdata.ParcelType
	ServerId               int64
	UniqueId               int64
	Level                  int32
	Exp                    int64
	Tier                   int32
	SlotIndex              int64
	BoundCharacterServerId int64
}
