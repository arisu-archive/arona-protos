package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type WeaponDB struct {
	ParcelBase
	Type                   flatdata.ParcelType
	UniqueId               int64
	Level                  int32
	Exp                    int64
	StarGrade              int32
	BoundCharacterServerId int64
}
