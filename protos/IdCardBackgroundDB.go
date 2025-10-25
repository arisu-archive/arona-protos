package protos

import (
	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type IdCardBackgroundDB struct {
	ParcelBase
	Type        flatdata.ParcelType
	ServerId    int64
	UniqueId    int64
	ParcelInfos []ParcelInfo
}
