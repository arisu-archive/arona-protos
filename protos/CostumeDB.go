package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type CostumeDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	BoundCharacterServerId int64 `json:",omitempty,omitzero"`
}
