package protos

import (
	"time"

	"github.com/arisu-archive/arona-flatbuffers/go/flatdata"
)

type EmblemDB struct {
	ParcelBase
	Type        flatdata.ParcelType
	UniqueId    int64
	ReceiveDate time.Time
	ParcelInfos []ParcelInfo
}
