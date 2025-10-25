package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type StickerDB struct {
	ParcelBase
	Type            flatdata.ParcelType
	StickerUniqueId int64
	ParcelInfos     []ParcelInfo
}
