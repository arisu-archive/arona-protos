package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type StickerDB struct {
	ParcelBase
	Type flatdata.ParcelType `json:",omitempty,omitzero"`
	StickerUniqueId int64 `json:",omitempty,omitzero"`
	ParcelInfos []ParcelInfo `json:",omitempty,omitzero"`
}
