package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CostumeDB struct {
	ParcelBase
	Type                   flatdata.ParcelType
	UniqueId               int64
	BoundCharacterServerId int64
}
